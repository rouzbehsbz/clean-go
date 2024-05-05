package postgres

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var instance *PostgresDB
var lock = &sync.Mutex{}

type PostgresDB struct {
	Db *gorm.DB
}

func newPostgresDB(connestionString string) (*PostgresDB, error) {
	p := new(PostgresDB)

	db, err := gorm.Open(postgres.Open(connestionString))

	if err != nil {
		return nil, err
	}

	p.Db = db

	return p, nil
}

func GetInstance(connestionString string) (*PostgresDB, error) {
	if instance == nil {
		lock.Lock()

		defer lock.Unlock()

		if instance == nil {
			ins, err := newPostgresDB(connestionString)

			if err != nil {
				return nil, err
			}

			instance = ins
		}
	}

	return instance, nil
}
