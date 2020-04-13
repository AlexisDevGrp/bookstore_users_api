package services

import (
	"fmt"
	"github.com/AlexisDevGrp/bookstore_users_api/domain/users"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
)

func CreateUser (user users.User) (*users.User, *mess.RestMsg) {



	fmt.Println("Try to create user : ", user)

	return &user, nil
}