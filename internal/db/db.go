package db

import (
"github.com/jmoiron/sqlx"
"github.com/cuelabs/sptfy/pkg/user"
_ "github.com/lib/pq"
"log"
)

type Store interface {
	createUser(user *user.SptfyUser) error
	readUser(id int) (*user.SptfyUser, error)
}

type DB struct {
	*sqlx.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		log.Print("Could not create a new database in NewDB()")
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Print("Could not ping database in NewDB()")
		return nil, err
	}
	return &DB{db}, nil
}
