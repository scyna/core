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
	TRACE_TASK         TraceType = 3
	TRACE_DOMAIN_EVENT TraceType = 4
)

type trace struct {
	ParentID  uint64
	ID        uint64
	Type      TraceType
	Time      time.Time
	Duration  uint64
	Path      string
	SessionID uint64
	Status    uint32
	Source    string
}

func CreateTrace(path string, type_ TraceType, parent ...uint64) *trace {
	var parent_ uint64 = 0
	if len(parent) == 1 {
		parent_ = parent[0]
	}
	return &trace{
		ParentID:  parent_,
		ID:        ID.Next(),
		Type:      type_,
		Time:      time.Now(),
		Path:      path,
		SessionID: Session.ID(),
		Source:    Module(),
	}
}

func (trace *trace) Record() {
	trace.Duration = uint64(time.Now().UnixNano() - trace.Time.UnixNano())
	EmitSignal(scyna_const.TRACE_CREATED_CHANNEL, &scyna_proto.TraceCreatedSignal{
		ID:        trace.ID,
		ParentID:  trace.ParentID,
		Type:      uint32(trace.Type),
		Time:      uint64(trace.Time.UnixMicro()),
		Duration:  trace.Duration,
		Path:      trace.Path,
		SessionID: trace.SessionID,
		Status:    trace.Status,
		Source:    trace.Source,
	})
}
