package users

import (
	"encoding/json"
	"fmt"
	"github.com/AlexisDevGrp/bookstore_users_api/domain/users"
	"github.com/AlexisDevGrp/bookstore_users_api/services"
	"github.com/AlexisDevGrp/bookstore_users_api/utils/mess"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	params:= mux.Vars(r)
	userStr := params["user_id"]
	userId, err := strconv.ParseInt(userStr, 10, 64)
	fmt.Println("userID: ", userId)
	if err != nil {
		msg := mess.BadReqErr("Invalid User Id. The user is mandatory for this type of request.")
		mess.SendMsg(w, *msg)
		return
	}
	user, errGet := services.GetUser(userId)
	if err != nil {
		msg := mess.ItemNotFound(errGet.Message)
		mess.SendMsg(w, *msg)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	obj, _ := json.Marshal(user);
	w.Write(obj)


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
	result, retMsg := services.CreateUser(user)
	if retMsg != nil {
		msg := mess.ItemNotCreated(retMsg.Message)
		mess.SendMsg(w, *msg)
		return
	}

	ResMsg := "User: " + strconv.FormatInt(result.Id, 10) + "has been created"
	msg := mess.ItemCreated(ResMsg)
	mess.SendMsg(w, *msg)
	return
}