package endpoint

import (
	"net/http"

	"github.com/api-skeleton/service/UserService"
)

func RegistrationEndpoint(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		UserService.UserRegistration(response, request)
		break
	}
}

func UserUpdateEndpoint(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "PUT":
		// UserService.UserProfileUpdate(response, request)
		break
	case "GET":
		// UserService.GetUserProfile(response, request)
		break
	}
}

func LoginEndpoint(response http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		// UserService.LoginService(response, request)
		break
	}
}
