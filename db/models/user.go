package models

import "casorder/utils/types"

// User data model
type User struct {
	ID           int64
	Username     string
	FullName     string
	PasswordHash string
}

// Serialize serializes user data
func (u *User) Serialize() types.JSON {
	return types.JSON{
		"id":           u.ID,
		"username":     u.Username,
		"display_name": u.FullName,
	}
}

func (u *User) Read(m types.JSON) {
	u.ID = int64(m["id"].(int64))
	u.Username = m["username"].(string)
	u.FullName = m["fullname"].(string)
}
