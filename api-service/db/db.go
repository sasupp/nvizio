package db

import (
	"github.com/gocql/gocql"
)

type DbConfig struct {
	Hosts []string
}

type Db struct {
	config  DbConfig
	session *gocql.Session
}

func NewDb(config DbConfig) *Db {
	return &Db{
		config: config,
	}
}

func (db *Db) Connect() (err error) {
	c := gocql.NewCluster(db.config.Hosts...)
	fallback := gocql.RoundRobinHostPolicy()
	c.PoolConfig.HostSelectionPolicy = gocql.TokenAwareHostPolicy(fallback)
	c.Consistency = gocql.One
	db.session, err = c.CreateSession()
	return err
}

func (db *Db) Close() {
	db.session.Close()
}
