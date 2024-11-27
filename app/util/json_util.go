package util

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func HandleRequest(writer http.ResponseWriter, dataFunc func() (interface{}, error)) {
	writer.Header().Add("Content-Type", "application/json")
	data, code, err := HandleResponse(dataFunc())
	writer.WriteHeader(code)
	if err != nil {
		log.Fatal(err)
	}
	_, err = fmt.Fprint(writer, string(data))
	if err != nil {
		log.Fatal(err)
	}
}

func HandleResponse(data interface{}, err error) (string, int, error) {
	var code = 200
	var message = "Success"
	if err != nil {
		code = 500
		message = err.Error()
	}
	serialisedData, er := Serialise(data, code, message)
	if er != nil {
		code = 500
	}
	return serialisedData, code, er

}

func Serialise(data interface{}, code int, message string) (string, error) {
	dataJson, err := json.Marshal(map[string]interface{}{"status": code, "data": data, "createdAt": time.Now(), "updatedAt": time.Now(), "message": message})
	return string(dataJson), err
}
