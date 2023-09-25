package scyna

import (
	"log"
	sync "sync"
	"time"

	scyna_const "github.com/scyna/core/const"
	scyna_proto "github.com/scyna/core/proto"
)

type session struct {
	id       uint64
	mutex    sync.Mutex
	sequence uint64
	quit     chan struct{}
}

func NewSession(id uint64) *session {
	ret := &session{
		id:       id,
		sequence: 1,
		quit:     make(chan struct{}),
	}

	ticker := time.NewTicker(10 * time.Minute)
	go func() {
		for {
			select {
			case <-ticker.C:
				EmitSignal(scyna_const.SESSION_UPDATE_CHANNEL, &scyna_proto.UpdateSessionSignal{ID: ret.id, Module: module})
			case <-ret.quit:
				ticker.Stop()
				return
			}
		}
	}()
	return ret
}

func (s *session) ID() uint64 {
	return s.id
}

func (s *session) NextSequence() uint64 {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.sequence++
	return s.sequence
}

func (s *session) release() {
	close(s.quit)
}

func (l *session) writeLog(level LogLevel, message string) {
	message = formatLog(message)
	log.Print(message)
	if l.id > 0 {
		AddLog(LogData{
			ID:       l.id,
			Sequence: Session.NextSequence(),
			Level:    level,
			Message:  message,
		})
	}
}

func (l *session) Info(messsage string) {
	l.writeLog(LOG_INFO, messsage)
}

func (l *session) Error(messsage string) {
	l.writeLog(LOG_ERROR, messsage)
}

func (l *session) Warning(messsage string) {
	l.writeLog(LOG_WARNING, messsage)
}

func (l *session) Debug(messsage string) {
	l.writeLog(LOG_DEBUG, messsage)
}

func (l *session) Fatal(messsage string) {
	l.writeLog(LOG_FATAL, messsage)
}
