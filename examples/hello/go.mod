module github.com/scyna/core/examples/hello

go 1.18

replace github.com/scyna/core => ../..

require (
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/scyna/core v0.0.0-00010101000000-000000000000
	github.com/scyna/core/example v0.0.0-20221206012752-ddfcaffcdd2e
	google.golang.org/protobuf v1.28.0
)

require (
	github.com/gocql/gocql v1.0.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	github.com/nats-io/nats.go v1.14.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/scylladb/go-reflectx v1.0.1 // indirect
	github.com/scylladb/gocqlx/v2 v2.7.0 // indirect
	golang.org/x/crypto v0.0.0-20220411220226-7b82a4e95df4 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)
