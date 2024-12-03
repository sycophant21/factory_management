package http

import (
	"encoding/json"
	logg "factory_management_go/app/log"
	"fmt"
	"net/http"
)

func HandleRequest[T any](writer http.ResponseWriter, dataFunc func() (T, error), dataConverterFunc func(data T) interface{}) {
	writer.Header().Add("Content-Type", "application/json")
	dataEntity, err := dataFunc()
	var data string
	var code int
	if err != nil {
		logg.Logger.Error(err.Error())
		data, code, err = HandleResponse(nil, err)
	} else {
		data, code, err = HandleResponse(dataConverterFunc(dataEntity), nil)
	}
	writer.WriteHeader(code)
	if err != nil {
		logg.Logger.Error(err.Error())
	}
	_, err = fmt.Fprint(writer, data)
	if err != nil {
		logg.Logger.Error(err.Error())
	}
}

func HandleResponse(data interface{}, err error) (string, int, error) {
	var code = 200
	if err != nil {
		code = 500
	}
	serialisedData, er := Serialise(data)
	if er != nil {
		code = 500
	}
	return serialisedData, code, er

}

func Serialise(data interface{}) (string, error) {
	dataJson, err := json.Marshal(data)
	return string(dataJson), err
}
