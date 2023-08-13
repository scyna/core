package scyna_setup

import (
	"github.com/gocql/gocql"
	scyna_const "github.com/scyna/core/const"
)

type module struct {
	code      string
	secret    string
	batch     *gocql.Batch
	consumers []string
}

func NewModule(code string, secret string) *module {
	ret := &module{
		code:   code,
		secret: secret,
		batch:  DB.Session.NewBatch(gocql.UnloggedBatch),
	}
	ret.batch.Query("INSERT INTO "+scyna_const.MODULE_TABLE+"(code, secret) VALUES(?,?)", code, secret)
	return ret
}

func (m *module) AddConsumer(consumer string) *module {
	m.consumers = append(m.consumers, consumer)
	return m
}

func (m *module) AddSetting(key, value string) *module {
	m.batch.Query("INSERT INTO "+scyna_const.SETTING_TABLE+"(module, key, value) VALUES(?,?,?)", m.code, key, value)
	return m
}

func (m *module) Build() {
	if err := DB.Session.ExecuteBatch(m.batch); err != nil {
		panic("Error in creating module:" + err.Error())
	}

	CreateStreamIfMissing(m.code)

	for _, e := range m.consumers {
		CreateConsumerIfMissing(m.code, e)
	}
}
