package UserService

import (
	"database/sql"
	"net/http"

	"github.com/api-skeleton/config"
	"github.com/api-skeleton/constanta"

	"github.com/api-skeleton/dao"
	"github.com/api-skeleton/dto/in"
	"github.com/api-skeleton/dto/out"
	"github.com/api-skeleton/model"
	"github.com/api-skeleton/utils"
)

func UserRegistration(res http.ResponseWriter, req *http.Request) (err error) {
	var (
		reqBody model.UserModel
		db      = config.Connect()
	)

	reqBody = inputStruct(utils.GetUserBody(req))

	err = dao.UserDAO.InsertUser(db, reqBody)
	if err != nil {
		out.ResponseOut(res, nil, false, constanta.CodeInternalServerErrorResponse, constanta.ErrorInternalDB)
		return
	}
	db.Close()
	out.ResponseOut(res, nil, true, constanta.CodeSuccessResponse, constanta.SuccessRegistrationData)
	return
}

func inputStruct(reqBody in.UserRequest) model.UserModel {
	reqBody.ValidationRegistration()
	return model.UserModel{
		ID:        sql.NullInt64{Int64: reqBody.Id},
		Username:  sql.NullString{String: reqBody.Username},
		Password:  sql.NullString{String: reqBody.Password},
		FirstName: sql.NullString{String: reqBody.FirstName},
		LastName:  sql.NullString{String: reqBody.LastName},
		Gender:    sql.NullString{String: reqBody.Gender},
		Telephone: sql.NullString{String: reqBody.Telephone},
		Email:     sql.NullString{String: reqBody.Email},
		Address:   sql.NullString{String: reqBody.Address},
	}
}
