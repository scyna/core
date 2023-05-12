package scyna_setup

import (
	"github.com/gocql/gocql"
	scyna_const "github.com/scyna/core/const"
)

type application struct {
	code  string
	auth  string
	batch *gocql.Batch
}

func NewApplication(code, auth string) *application {
	ret := &application{code: code, auth: auth, batch: DB.NewBatch(gocql.UnloggedBatch)}
	ret.batch.Query("INSERT INTO "+scyna_const.APPLICATION_TABLE+"(code, auth_url) VALUES(?,?)", code, auth)
	return ret
}

func (c *application) UseEndpoint(url string) *application {
	c.batch.Query("INSERT INTO "+scyna_const.APP_USE_ENDPOINT_TABLE+"(application, url) VALUES(?,?)", c.code, url)
	return c
}

func (c *application) Build() {
	if err := DB.Session.ExecuteBatch(c.batch); err != nil {
		panic("Error in creating client:" + err.Error())
	}
}
