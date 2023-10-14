package scyna_setup

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"
	scyna "github.com/scyna/core"
)

type db struct {
	Session *gocql.Session
}

func NewDB(host []string, username string, password string, location string) *db {
	cluster := gocql.NewCluster(host...)
	cluster.Authenticator = gocql.PasswordAuthenticator{Username: username, Password: password}
	cluster.PoolConfig.HostSelectionPolicy = gocql.DCAwareRoundRobinPolicy(location)
	cluster.ConnectTimeout = time.Second * 10
	cluster.DisableInitialHostLookup = true
	cluster.Consistency = gocql.Quorum

	session, err := cluster.CreateSession()
	if err != nil {
		panic(fmt.Sprintf("Can not create session: Host = %s, Error = %s ", host, err.Error()))
	}
	return &db{Session: session}
}

func (db *db) QueryOne(query string, args ...interface{}) *gocql.Query {
	return db.Session.Query(addLimitOne(query), args...).
		WithContext(context.Background()).
		Consistency(gocql.One)
}

func (db *db) QueryMany(query string, args ...interface{}) gocql.Scanner {
	return db.Session.Query(query, args...).
		WithContext(context.Background()).
		Iter().Scanner()
}

func (db *db) Execute(query string, args ...interface{}) error {
	return db.Session.Query(query, args...).Exec()
}

func (db *db) Apply(query string, args ...interface{}) (bool, error) {
	dest := make(map[string]interface{})
	return db.Session.Query(query, args...).MapScanCAS(dest)
}

func (db *db) AssureExists(query string, args ...interface{}) scyna.Error {
	scanner := db.Session.Query(addLimitOne(query), args...).
		WithContext(context.Background()).
		Consistency(gocql.One).Iter().Scanner()
	if !scanner.Next() {
		return scyna.OBJECT_NOT_FOUND
	}
	return nil
}

func (db *db) AssureNotExists(query string, args ...interface{}) scyna.Error {
	scanner := db.Session.Query(addLimitOne(query), args...).
		WithContext(context.Background()).
		Consistency(gocql.One).Iter().Scanner()
	if scanner.Next() {
		return scyna.OBJECT_EXISTS
	}
	return nil
}

func (db *db) Close() {
	db.Session.Close()
}

func addLimitOne(query string) string {
	if strings.Contains(strings.ToUpper(query), "LIMIT") {
		return query
	}
	return query + " LIMIT 1"
}
