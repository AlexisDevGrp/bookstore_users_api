package services

import (
	"fmt"
	"github.com/AlexisDevGrp/bookstore_users_api/domain/users"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
)

func CreateUser (user users.User) (*users.User, *mess.RestMsg) {
    if err := user.Validate(); err != nil {
    	return nil, err
	}
	if err := user.Save(); err != nil {
		return nil, err
	}
	return &user, nil
}
func GetUser (userId int64) (*users.User, *mess.RestMsg) {
	res := &users.User{Id: userId}

	if err := res.Get(); err != nil {
		return nil, err
	}
	fmt.Println("Find: res")
	return res, nil
}