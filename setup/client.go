package scyna_setup

import "github.com/gocql/gocql"

type client struct {
	id     string
	secret string
	batch  *gocql.Batch
}

func NewClient(id, secret string) *client {
	ret := &client{id: id, secret: secret, batch: DB.NewBatch(gocql.UnloggedBatch)}
	ret.batch.Query("INSERT INTO scyna.client(id, secret) VALUES(?,?)", id, secret)
	return ret
}

func (c *client) UseEndpoint(url string) *client {
	c.batch.Query("INSERT INTO scyna.client_use_endpoint(client, url) VALUES(?,?)", c.id, url)
	return c
}

func (c *client) Build() {
	if err := DB.Session.ExecuteBatch(c.batch); err != nil {
		panic("Error in creating client:" + err.Error())
	}
}
