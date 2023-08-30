protoc -I=. --go_out=generated scyna.proto
protoc -I=. --go_out=generated error.proto
protoc -I=. --go_out=generated task.proto
protoc -I=. --go_out=generated session.proto
protoc -I=. --go_out=generated setting.proto
protoc -I=. --go_out=generated id.proto
protoc -I=. --go_out=generated trace.proto