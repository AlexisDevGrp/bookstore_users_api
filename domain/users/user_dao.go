package users

import (
	"fmt"
	"github.com/AlexisDevGrp/bookstore_users_api/datasources/ora/users_db"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
)
const(
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, data_created) VALUES(?,?,?,?);"
)
var (
	usersDB = make(map[int64]*User)
)
func (user *User) Get() *mess.RestMsg {
	fmt.Println("Trying Ping on DB again")
	fmt.Println("It works")
	res := usersDB[user.Id]
	if res == nil {
		return mess.NotFound(fmt.Sprintf("User %d not found", user.Id))
	}
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.Email = res.Email
	user.DateCreated = res.DateCreated

	return nil
}
func (user *User) Save() *mess.RestMsg{
	users_db.Init()
	err := users_db.SaveUser(queryInsertUser, user); if err != nil {
		return err
	}
	return nil
}
