package out

import (
	"encoding/json"
	"net/http"

	"github.com/api-skeleton/utils"
)

type APIResponse struct {
	API APIMessage `json:"API"`
}

type APIMessage struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func (ar APIResponse) String() string {
	return utils.StructToJSON(ar)
}

func ResponseOut(response http.ResponseWriter, data interface{}, success bool, code int, message string) {
	response.Header().Set("Content-type", "application/json")
	var apiResponse APIResponse
	apiResponse.API.Success = success
	apiResponse.API.Code = code
	apiResponse.API.Message = message
	apiResponse.API.Content = data
	response.WriteHeader(code)
	json.NewEncoder(response).Encode(apiResponse)
}
