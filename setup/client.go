package scyna_setup

import (
	"github.com/gocql/gocql"
	scyna_const "github.com/scyna/core/const"
)

type client struct {
	id     string
	secret string
	batch  *gocql.Batch
}

func NewClient(id, secret string) *client {
	ret := &client{id: id, secret: secret, batch: DB.Session.NewBatch(gocql.UnloggedBatch)}
	ret.batch.Query("INSERT INTO "+scyna_const.CLIENT_TABLE+"(id, secret) VALUES(?,?)", id, secret)
	return ret
}

func (c *client) UseEndpoint(url string) *client {
	c.batch.Query("INSERT INTO "+scyna_const.CLIENT_USE_ENDPOINT_TABLE+"(client, url) VALUES(?,?)", c.id, url)
	return c
}

func (c *client) Build() {
	if err := DB.Session.ExecuteBatch(c.batch); err != nil {
		panic("Error in creating client:" + err.Error())
	}
}
