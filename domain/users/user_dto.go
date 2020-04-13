package users

import (
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
	"strings"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name""`
	LastName string `json:"last_name""`
	Email string `json:"email""`
	DateCreated string `json:"date_created""`
}

func (user *User) Validate() *mess.RestMsg {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return mess.BadReqErr("Invalid Email Address")
	}
	return nil
}