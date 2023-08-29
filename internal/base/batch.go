package base

import "github.com/gocql/gocql"

type batch struct {
	batch *gocql.Batch
	db    *DB
}

func (b *batch) Add(query string, args ...interface{}) *batch {
	b.batch.Query(query, args...)
	return b
}

func (b *batch) Execute() error {
	return b.db.Session.ExecuteBatch(b.batch)
}
