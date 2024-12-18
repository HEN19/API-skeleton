package dao

import (
	"database/sql"
	"fmt"

	"github.com/api-skeleton/model"
)

type userDAO struct {
	AbstractDAO
}

var UserDAO = userDAO{}.New()

func (input userDAO) New() (output userDAO) {
	output.TableName = "users"
	output.FileName = "UserDAO.go"
	return
}

func (input userDAO) InsertUser(db *sql.DB, inputStruct model.UserModel) (err error) {
	var (
		query string
	)

	query = fmt.Sprintf(
		`
			INSERT INTO %s
				(username, password, first_name, last_name, gender, telephone, email, address)
			VALUES (?,?,?,?,?,?,?,?)
		`, input.TableName,
	)

	params := []interface{}{
		inputStruct.Username, inputStruct.Password, inputStruct.FirstName,
		inputStruct.LastName, inputStruct.Gender, inputStruct.Telephone,
		inputStruct.Email, inputStruct.Address,
	}

	_, err = db.Exec(query, params...)
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
