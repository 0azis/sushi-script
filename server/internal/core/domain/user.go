package domain

type User struct {
	ID      int    `json:"ID"`
	Phone   string `json:"phone"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (u *User) Validate() bool {
	return u.Phone != ""
}
