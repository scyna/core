protoc -I=. --go_out=generated scyna.proto
protoc -I=. --go_out=generated error.proto
protoc -I=. --go_out=generated task.proto
protoc -I=. --go_out=generated engine.proto

