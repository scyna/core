package scyna

import (
	"log"
	"math"
	"reflect"
	"time"

	"github.com/gocql/gocql"
	"google.golang.org/protobuf/proto"
)

type Activity struct {
	Type    string
	Event   proto.Message
	Version int64
	Time    time.Time
}

type EventStore[T proto.Message] struct {
	Table       string
	projections map[string]IProjection
}

func NewEventStore[T proto.Message](table string) *EventStore[T] {
	return &EventStore[T]{
		Table:       table,
		projections: make(map[string]IProjection),
	}
}

func (e *EventStore[T]) registerProjection(projection IProjection) {
	e.projections[projection.EventName()] = projection
}

func (e *EventStore[T]) ReadModel(id any) (*Model[T], Error) {
	var version int64
	var data []byte

	if err := DB.QueryOne("SELECT version, data FROM "+e.Table+
		" WHERE id=? LIMIT 1", id).Scan(&version, &data); err != nil {
		if err == gocql.ErrNotFound {
			return nil, OBJECT_NOT_FOUND
		}
		return nil, SERVER_ERROR
	}

	var ret T
	if err := proto.Unmarshal(data, ret); err != nil {
		return nil, BAD_DATA
	}

	return &Model[T]{
		ID:      id,
		Version: version,
		Data:    ret,
		store:   e,
	}, nil
}

func (e *EventStore[T]) CreateModel(id any) (*Model[T], Error) {
	var version int64
	var data []byte

	if err := DB.QueryOne("SELECT version, data FROM "+e.Table+
		" WHERE id=? LIMIT 1", id).Scan(&version, &data); err != nil {
		if err == gocql.ErrNotFound {
			var ret T
			return &Model[T]{
				ID:      id,
				Version: 0,
				Data:    ret,
				store:   e,
			}, nil
		}
		return nil, SERVER_ERROR
	}
	return nil, OBJECT_EXISTS
}

func (e *EventStore[T]) updateWriteModel(model *Model[T], event proto.Message) Error {
	model.Version++

	eventData, err := proto.Marshal(event)
	if err != nil {
		return BAD_DATA
	}

	modelData, err := proto.Marshal(model.Data)
	if err != nil {
		return BAD_DATA
	}

	if applied, err := DB.Apply("INSERT INTO "+e.Table+
		" (id, type, event, data, version, created, state) VALUES (?, ?, ?, ?, ?, ?, ?) IF NOT EXISTS",
		model.ID, reflect.TypeOf(event).Elem().Name(), eventData, modelData, model.Version, time.Now(), 0); applied {
		return nil
	} else {
		if err != nil {
			log.Print("updateWriteModel:", err)
		}
	}

	return COMMAND_NOT_COMPLETED
}

func (e *EventStore[T]) updateReadModel(id any) {
	version := e.getLastSynced(id)
	if version == -1 {
		return
	}
	version++
	for e.doSync(id, version) {
		version++
	}
}

func (e *EventStore[T]) getLastSynced(id any) int64 {
	var version int64
	if err := DB.QueryOne("SELECT version FROM "+e.Table+
		" WHERE id=? AND state=? LIMIT 1", id, 2).Scan(&version); err != nil {
		if err == gocql.ErrNotFound {
			return 0
		}
		return -1
	}
	return version
}

func (e *EventStore[T]) tryToLock(id any, version int64) bool {
	if applied, _ := DB.Apply("UPDATE "+e.Table+" SET locked=?, state=? WHERE id=? AND version=? IF state=?",
		time.Now(), 1, id, version, 0); applied {
		return true
	}
	return false
}

func (e *EventStore[T]) lockLongLockingRow(id any, version int64) bool {
	var locked time.Time
	if err := DB.QueryOne("SELECT locked FROM "+e.Table+" WHERE id=? AND version=?",
		id, version).Scan(&locked); err != nil {
		return false
	}

	if locked.Add(time.Second * 5).Before(time.Now()) {
		if applied, _ := DB.Apply("UPDATE "+e.Table+" SET state=? WHERE id=? AND version=? IF state=?",
			0, id, version, 1); applied {
			return true
		}
	}
	return false
}

func (e *EventStore[T]) doSync(id any, version int64) bool {
	if !e.tryToLock(id, version) {
		if !e.lockLongLockingRow(id, version) {
			return false
		}
	}

	if !e.syncRow(id, version) {
		return false
	}

	if !e.markSynced(id, version) {
		return false
	}
	return true
}

func (e *EventStore[T]) syncRow(id any, version int64) bool {
	var type_ string
	var data []byte
	var event []byte

	if err := DB.QueryOne("SELECT type,data,event FROM "+e.Table+
		" WHERE id=? AND version=? LIMIT 1", id, version).Scan(&type_, &data, &event); err != nil {
		if err == gocql.ErrNotFound {
			return false
		}
	}

	p, ok := e.projections[type_]

	if !ok {
		log.Print("No projection for type=", type_)
		return false
	}

	p.Update(data, event)

	return true
}

func (e *EventStore[T]) markSynced(id any, version int64) bool {
	if err := DB.Execute("UPDATE "+e.Table+
		" SET state=? WHERE id=? AND version=?", 2, id, version); err != nil {
		log.Print("markSynced:", err)
		return false
	}
	return true
}

func (e *EventStore[T]) parseEvent(type_ string, data []byte) proto.Message {
	p, ok := e.projections[type_]

	if !ok {
		return nil
	}

	return p.ParseEvent(data)
}

func (e *EventStore[T]) ListActivity(id any, position int64, count int32) []Activity {
	if position == 0 {
		position = math.MaxInt64
	}
	if count == 0 {
		count = 50
	}
	if count > 100 {
		count = 100
	}

	var version int64
	var type_ string
	var event []byte
	var created time.Time

	rs := DB.QueryMany("SELECT version,type,event,created FROM "+e.Table+
		" WHERE id=? AND version<? LIMIT ?", id, position, count)

	var ret []Activity
	for rs.Next() {
		if err := rs.Scan(&version, &type_, &event, &created); err == nil {
			if data := e.parseEvent(type_, event); data != nil {
				ret = append(ret, Activity{
					Version: version,
					Type:    type_,
					Event:   data,
					Time:    created,
				})
			}
		}
	}
	return ret
}
