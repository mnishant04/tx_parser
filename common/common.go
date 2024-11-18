package common

import (
	"encoding/json"
	"log"
	"net/http"
)

type Response struct {
	Msg        string      `json:"msg"`
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
}

func SendResponse(w http.ResponseWriter, data any, msg string, statusCode int) {
	var jsonBytes []byte
	var err error
	w.WriteHeader(statusCode)
	if data != nil {
		jsonBytes, err = json.Marshal(data)
		if err != nil {
			log.Printf("Error marshalling data %s", err)
		}
	}
	w.Write(jsonBytes)
}
