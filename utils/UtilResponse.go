package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/api-skeleton/dto/in"
	"github.com/gorilla/mux"
)

func StructToJSON(input interface{}) (output string) {
	b, err := json.Marshal(input)
	if err != nil {
		fmt.Println(err)
		output = ""
		return
	}
	output = string(b)
	return
}

func readBody(request *http.Request) (output string) {
	byteBody, err := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	if err != nil {
		return ""
	}
	return string(byteBody)
}

func GetUserBody(request *http.Request) (bodyRequest in.UserRequest) {
	jsonString := readBody(request)
	json.Unmarshal([]byte(jsonString), &bodyRequest)
	return
}

func ReadParam(request *http.Request) (id int64, err error) {
	strId, ok := mux.Vars(request)["Id"]
	idParam, errConvert := strconv.Atoi(strId)
	id = int64(idParam)

	if !ok || errConvert != nil {
		err = errConvert
		return
	}
	return
}
