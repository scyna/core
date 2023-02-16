package scyna_setup

import "github.com/gocql/gocql"

type application struct {
	code  string
	auth  string
	batch *gocql.Batch
}

func NewApplication(code, auth string) *application {
	ret := &application{code: code, auth: auth, batch: DB.NewBatch(gocql.UnloggedBatch)}
	ret.batch.Query("INSERT INTO scyna.application(code, auth_url) VALUES(?,?)", code, auth)
	return ret
}

func (c *application) UseEndpoint(url string) *application {
	c.batch.Query("INSERT INTO scyna.application_use_endpoint(application, url) VALUES(?,?)", c.code, url)
	return c
}

func (c *application) Build() {
	if err := DB.Session.ExecuteBatch(c.batch); err != nil {
		panic("Error in creating client:" + err.Error())
	}
}
