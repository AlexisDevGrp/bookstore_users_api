package users

import (
	"fmt"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/date_utils"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
)

var (
	usersDB = make(map[int64]*User)
)
func (user *User) Get() *mess.RestMsg {
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
	current := usersDB[user.Id]
	if current != nil {
		return mess.BadReqErr(fmt.Sprintf("User %d already exists", user.Id))
	}
	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}
