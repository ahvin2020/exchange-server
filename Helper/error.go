package helper

import (
	"encoding/json"
	"net/http"
)

type MsgError struct {
	Message string `json:"message"`
	Type    int16  `json:"type"`
}

// type of errors
const ERR_SYSTEM int16 = 100 // system error, e.g. db syntax error
const ERR_AUTH int16 = 200   // authentication error, e.g. wrong password

func ReturnErr(w http.ResponseWriter, msg string, errType int16) {
	msgError := MsgError{
		Message: msg,
		Type:    errType,
	}

	data := struct {
		Error MsgError `json:"Error"`
	}{
		Error: msgError,
	}

	json.NewEncoder(w).Encode(&data)
}
