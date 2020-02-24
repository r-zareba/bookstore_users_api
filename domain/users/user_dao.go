// The only place to execute Database commands

package users

import (
	"fmt"
	"github.com/r-zareba/bookstore_users_api/datasources/mysql/users_db"
	"github.com/r-zareba/bookstore_users_api/utils/date_utils"
	"github.com/r-zareba/bookstore_users_api/utils/errors"
	"github.com/r-zareba/bookstore_users_api/utils/mysql_utils"
)

const (
	insertUserQuery   = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	getUserQuery      = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	updateQuery       = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	deleteQuery       = "DELETE FROM users WHERE id=?;"
	findByStatusQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?"
)

func (user *User) GetFromDB() *errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(getUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)
	// Scan - automatically fill the object attributes
	getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email,
		&user.DateCreated, &user.Status, &user.Password)
	if getErr != nil {
		return mysql_utils.GetMySQLError(getErr)
	}
	return nil
}

func (user *User) DeleteFromDB() *errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(deleteQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	_, deleteErr := statement.Exec(user.Id)
	if deleteErr != nil {
		return mysql_utils.GetMySQLError(deleteErr)
	}
	return nil
}

func (user *User) FindByStatusInDB(status string) ([]User, *errors.RestError) {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(findByStatusQuery)
	if err != nil {
		return nil, errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	rows, err := statement.Query(status)
	if err != nil {
		return nil, mysql_utils.GetMySQLError(err)
	}
	defer rows.Close()

	results := make([]User, 0)
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
		if err != nil {
			return nil, mysql_utils.GetMySQLError(err)
		}
		results = append(results, user)
	}

	if len(results) == 0 {
		return nil, errors.NotFoundError(fmt.Sprintf("Users with status: %s not found", status))
	}
	return results, nil
}

func (user *User) SaveToDB() *errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(insertUserQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	user.DateCreated = date_utils.GetNowTime()
	insertResult, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated,
		user.Status, user.Password)

	if saveErr != nil {
		return mysql_utils.GetMySQLError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.InternalServerError("Cannot get last inserted ID")
	}

	user.Id = userId
	return nil
}

func (user *User) UpdateInDB() *errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(updateQuery)
	if err != nil {
		return errors.InternalServerError(err.Error())
	}
	defer statement.Close()

	_, updateErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.Id)
	if updateErr != nil {
		return mysql_utils.GetMySQLError(updateErr)
	}
	return nil
}
