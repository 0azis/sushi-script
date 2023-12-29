package store

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type Store struct {
	DB *sqlx.DB
}

func (store *Store) Connect() {
	db, err := sqlx.Connect("mysql", "test:test123@(localhost:3333)/sushi")
	if err != nil {
		log.Println(err)
	}

	store.DB = db
}

func (store *Store) Close() {
	store.DB.Close()
}

func InitDB() *Store {
	return &Store{}
}
