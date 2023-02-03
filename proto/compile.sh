protoc -I=. --go_out=generated scyna.proto
protoc -I=. --go_out=generated error.proto
protoc -I=. --go_out=generated scheduler.proto

