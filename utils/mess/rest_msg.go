package mess

import (
	"encoding/json"
	"net/http"
)

type RestMsg struct {
	Message string `json:"message"`
	Status int `json: "code"`
	Error string `json: error`
}

func SendMsg(w http.ResponseWriter, msg RestMsg) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(msg.Status)
	obj, _ := json.Marshal(msg);
	w.Write(obj)
}
func BadReqErr(msg string) *RestMsg {
	return &RestMsg {
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "Bad Request",
	}
}
func BodyJsonErr(msg string) *RestMsg {
	return &RestMsg {
		Message: msg,
		Status:  http.StatusInternalServerError,
		Error:   "Body Json Problem",
	}
}
func NotFound(msg string) *RestMsg {
	return &RestMsg {
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "Not Found",
	}
}
func ItemCreated(msg string) *RestMsg {
	return &RestMsg {
		Message: msg,
		Status:  http.StatusCreated,
		Error:   "Item Created",
	}
}
func ItemNotCreated(msg string) *RestMsg {
	return &RestMsg {
		Message: msg,
		Status:  http.StatusBadRequest,
		Error:   "Item Not Created",
	}
}
func ItemNotFound(msg string) *RestMsg {
	return &RestMsg {
		Message: msg,
		Status:  http.StatusNotFound,
		Error:   "Item Not Found",
	}
}
