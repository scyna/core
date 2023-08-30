package scyna

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	scyna_const "github.com/scyna/core/const"
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
				if err := DB.Execute("INSERT INTO "+scyna_const.LOG_TABLE+
					" (source, day, time, seq, level, message) VALUES (?, ?, ?, ?, ?, ?)",
					l.ID, scyna_utils.GetDayByTime(time_), time_, l.Sequence, l.Level, l.Message); err != nil {
					log.Println("saveSessionLog: " + err.Error())
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
					Time:  uint64(time_),
					ID:    l.ID,
					Level: uint32(l.Level),
					Text:  l.Message,
					SEQ:   l.Sequence,
				}
				EmitSignal(scyna_const.LOG_CREATED_CHANNEL, &event)
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
