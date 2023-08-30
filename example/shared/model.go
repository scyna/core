package shared

import (
	scyna "github.com/scyna/core"
	PROTO "github.com/scyna/core/example/proto/generated"
)

var Store scyna.EventStore[*PROTO.AccountModel] = scyna.NewEventStore[*PROTO.AccountModel]("account")
