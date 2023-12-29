package store

import (
	"fmt"
	"sushi/internal/core/domain"
)

func (store *Store) Insert(user domain.User) (int, error) {
	var insertedID int

	tx, err := store.DB.Begin()
	if err != nil {
		return insertedID, err
	}

	_, err = tx.Exec(`insert into sushi_users (phone, name, address) values (?, ?, ?);`, user.Phone, user.Name, `{}`)
	if err != nil {
		return insertedID, err
	}
	rows, err := tx.Query("select LAST_INSERT_ID();")
	if err != nil {
		return insertedID, err
	}
	tx.Commit()

	for rows.Next() {
		rows.Scan(&insertedID)
	}

	return insertedID, nil
}

func (store *Store) Select(key string, value any) (domain.User, error) {
	var resultUser domain.User

	err := store.DB.Get(&resultUser, fmt.Sprintf("select * from sushi_users where %s = %v", key, value))

	return resultUser, err
}

func (store *Store) Update(user domain.User) error {
	_, err := store.DB.Query(`update sushi_users set phone = ?, name = ?, address = ? where id = ?`, user.Phone, user.Name, user.Address, user.ID)
	return err
}
