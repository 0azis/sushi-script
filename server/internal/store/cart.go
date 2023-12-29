package store

func (store *Store) Add(userID, productID int) error {
	_, err := store.DB.Query(`insert into sushi_cart (userid, productid) values (?, ?)`, userID, productID)

	return err
}

func (store *Store) Delete(userID, productID int) error {
	_, err := store.DB.Query(`delete from sushi_cart where userid = ? and productid = ?`, userID, productID)

	return err
}
