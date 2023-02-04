package scyna_setup

import (
	"github.com/gocql/gocql"
)

type module struct {
	code   string
	secret string
	batch  *gocql.Batch
	events []string
}

func NewModule(code string, secret string) *module {
	ret := &module{
		code:   code,
		secret: secret,
		batch:  DB.NewBatch(gocql.UnloggedBatch),
	}
	ret.batch.Query("INSERT INTO scyna.module(code, secret) VALUES(?,?)", code, secret)
	return ret
}

func (m *module) AddEventChannel(sender string) *module {
	m.events = append(m.events, sender)
	return m
}

func (m *module) AddSetting(key, value string) *module {
	m.batch.Query("INSERT INTO scyna.setting(module, key, value) VALUES(?,?,?)", m.code, key, value)
	return m
}

func (m *module) Build() {
	if err := DB.Session.ExecuteBatch(m.batch); err != nil {
		panic("Error in creating module:" + err.Error())
	}

	CreateStreamIfMissing(m.code)

	for _, e := range m.events {
		CreateConsumerIfMissing(e, m.code)
	}
}
