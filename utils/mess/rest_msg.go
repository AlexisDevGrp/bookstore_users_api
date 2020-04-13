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
		Message: "Body Reception problem",
		Status:  http.StatusBadRequest,
		Error:   msg,
	}
}
func BodyJsonErr(msg string) *RestMsg {
	return &RestMsg {
		Message: "Body Json Problem",
		Status:  http.StatusInternalServerError,
		Error:   msg,
	}
}
