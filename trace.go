package scyna

import (
	"time"

	scyna_const "github.com/scyna/core/const"
	scyna_proto "github.com/scyna/core/proto/generated"
)

type TraceType uint32

const (
	TRACE_ENDPOINT     TraceType = 1
	TRACE_EVENT        TraceType = 2
	TRACE_SYNC         TraceType = 4
	TRACE_TASK         TraceType = 5
	TRACE_DOMAIN_EVENT TraceType = 6
)

type Trace struct {
	ParentID    uint64    `db:"parent_id"`
	ID          uint64    `db:"id"`
	Type        TraceType `db:"type"`
	Time        time.Time `db:"time"`
	Duration    uint64    `db:"duration"`
	Path        string    `db:"path"`
	Source      string    `db:"source"`
	SessionID   uint64    `db:"session_id"`
	Status      int32     `db:"status"`
	RequestBody string
}

func (trace *Trace) Record() {
	trace.Duration = uint64(time.Now().UnixNano() - trace.Time.UnixNano())
	EmitSignal(scyna_const.TRACE_CREATED_CHANNEL, &scyna_proto.TraceCreatedSignal{
		ID:        trace.ID,
		ParentID:  trace.ParentID,
		Type:      uint32(trace.Type),
		Time:      uint64(trace.Time.UnixMicro()),
		Duration:  trace.Duration,
		Path:      trace.Path,
		Source:    trace.Source,
		SessionID: trace.SessionID,
		Status:    trace.Status,
	})
}
