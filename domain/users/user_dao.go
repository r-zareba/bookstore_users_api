// The only place to execute Database commands

package users

import (
	"fmt"
	"github.com/r-zareba/bookstore_users_api/datasources/mysql/users_db"
	"github.com/r-zareba/bookstore_users_api/utils/date_utils"
	"github.com/r-zareba/bookstore_users_api/utils/errors"
	"strings"
)

const (
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
	getUserQuery    = "SELECT id, first_name, last_name, email, date_created FROM users WHERE id=?"

	uniqueEmailError = "email_UNIQUE"
	noRowsError      = "no rows in result set"
)

var usersDB = make(map[int64]*User)

// TODO Refactor? GetFromDB, FillFromDB
func (user *User) Get() *errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(getUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)
	// Scan - automatically fill the object attributes
	err = result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated)
	if err != nil {
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NotFoundError(fmt.Sprintf("User id: %d not found in database", user.Id))
		}
		return errors.InternalServerError(fmt.Sprintf("Error while getting user (id:%d)", user.Id))
	}
	return nil
}

func (user *User) Save() *errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(insertUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	user.DateCreated = date_utils.GetNowTime()
	insertResult, err := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated)

	if err != nil {
		if strings.Contains(err.Error(), uniqueEmailError) {
			return errors.BadRequestError(fmt.Sprintf("User with email: %s already exists", user.Email))
		}
		return errors.InternalServerError(fmt.Sprintf("Cannot insert user, %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError("Cannot get last inserted ID")
	}

	user.Id = userId
	return nil
}
