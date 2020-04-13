package users

import (
	"encoding/json"
	"fmt"
	"github.com/AlexisDevGrp/bookstore_users_api/domain/users"
	"github.com/AlexisDevGrp/bookstore_users_api/services"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
	"io/ioutil"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	fmt.Println(user)
	bytes, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(bytes))
	fmt.Println(err)
	return

}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var user users.User
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg := mess.BadReqErr(err.Error())
		mess.SendMsg(w, *msg)
		return
	}
	if err := json.Unmarshal(b, &user); err != nil {
		msg := mess.BodyJsonErr(err.Error())
		mess.SendMsg(w, *msg)
		return
	}

	result, saveMsg := services.CreateUser(user)
	if saveMsg != nil {
		restErr := mess.RestMsg{
			Message: "CreateUser not possible",
			Status:  http.StatusBadRequest,
			Error:   saveMsg.Message,
		}
		mess.SendMsg(w, restErr)
		fmt.Println("user:", result, " has been created")
		return
	}
}