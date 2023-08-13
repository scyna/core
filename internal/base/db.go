package base

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

type DB struct {
	Session *gocql.Session
}

var (
	OBJECT_NOT_FOUND = NewError(8, "Object Not Found")
	OBJECT_EXISTS    = NewError(9, "Object Exists")
)

func NewDB(host []string, username string, password string, location string) *DB {
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
	return &DB{Session: session}
}

func (db *DB) QueryOne(query string, args ...interface{}) *gocql.Query {
	return db.Session.Query(addLimitOne(query), args...).
		WithContext(context.Background()).
		Consistency(gocql.One)
}

func (db *DB) QueryMany(query string, args ...interface{}) gocql.Scanner {
	return db.Session.Query(query, args...).
		WithContext(context.Background()).
		Iter().Scanner()
}

func (db *DB) Execute(query string, args ...interface{}) error {
	return db.Session.Query(query, args...).Exec()
}

func (db *DB) AssureExists(query string, args ...interface{}) *Error {
	scanner := db.Session.Query(addLimitOne(query), args...).
		WithContext(context.Background()).
		Consistency(gocql.One).Iter().Scanner()
	if !scanner.Next() {
		return OBJECT_NOT_FOUND
	}
	return nil
}

func (db *DB) AssureNotExists(query string, args ...interface{}) *Error {
	scanner := db.Session.Query(addLimitOne(query), args...).
		WithContext(context.Background()).
		Consistency(gocql.One).Iter().Scanner()
	if scanner.Next() {
		return OBJECT_EXISTS
	}
	return nil
}

func (db *DB) Close() {
	db.Session.Close()
}

func addLimitOne(query string) string {
	if strings.Contains(strings.ToUpper(query), "LIMIT") {
		return query
	}
	return query + " LIMIT 1"
}
