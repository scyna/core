package scyna

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/scylladb/gocqlx/v2/qb"
	scyna_proto "github.com/scyna/core/proto/generated"
	scyna_utils "github.com/scyna/core/utils"
)

type LogLevel int

const (
	LOG_INFO    LogLevel = 1
	LOG_ERROR   LogLevel = 2
	LOG_WARNING LogLevel = 3
	LOG_DEBUG   LogLevel = 4
	LOG_FATAL   LogLevel = 5
)

type Logger interface {
	Info(messsage string)
	Error(messsage string)
	Warning(messsage string)
	Debug(messsage string)
	Fatal(messsage string)
}

type LogData struct {
	Level    LogLevel
	Message  string
	ID       uint64
	Sequence uint64
	Session  bool
}

type TraceLogger struct {
	TraceID uint64
}

var logQueue chan LogData

func UseDirectLog(count int) {
	logQueue = make(chan LogData)

	for i := 0; i < count; i++ {
		go func() {
			for l := range logQueue {
				time_ := time.Now()
				if l.Session {
					if err := qb.Insert("scyna.session_log").
						Columns("session_id", "day", "time", "seq", "level", "message").
						Query(DB).
						Bind(l.ID, scyna_utils.GetDayByTime(time_), time_, l.Sequence, l.Level, l.Message).
						ExecRelease(); err != nil {
						log.Println("saveSessionLog: " + err.Error())
					}
				} else {
					if err := qb.Insert("scyna.log").
						Columns("trace_id", "time", "seq", "level", "message").
						Query(DB).
						Bind(l.ID, time_, l.Sequence, l.Level, l.Message).
						ExecRelease(); err != nil {
						log.Println("saveServiceLog: " + err.Error())
					}
				}
			}
		}()
	}
}

func UseRemoteLog(count int) {
	logQueue = make(chan LogData)

	for i := 0; i < count; i++ {
		go func() {
			for l := range logQueue {
				time_ := time.Now().UnixMicro()
				event := scyna_proto.LogCreatedSignal{
					Time:    uint64(time_),
					ID:      l.ID,
					Level:   uint32(l.Level),
					Text:    l.Message,
					Session: l.Session,
					SEQ:     l.Sequence,
				}
				EmitSignal(scyna_proto.LOG_CREATED_CHANNEL, &event)
			}
		}()
	}
}

func AddLog(data LogData) {
	if logQueue != nil {
		logQueue <- data
	}
}

func releaseLog() {
	if logQueue != nil {
		close(logQueue)
	}
}

func formatLog(message string) string {
	pc, file, line, ok := runtime.Caller(3)
	if !ok {
		return fmt.Sprintf("[?:0 - ?] %s", message)
	}
	path := strings.Split(file, "/")
	filename := path[len(path)-1]

	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return fmt.Sprintf("[%s:%d - ?] %s", filename, line, message)
	}
	fPath := strings.Split(fn.Name(), "/")
	funcName := fPath[len(fPath)-1]
	return fmt.Sprintf("[%s:%d - %s] %s", filename, line, funcName, message)
}
