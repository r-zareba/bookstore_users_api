// The only place to execute Database commands

package users

import (
	"fmt"
	"github.com/r-zareba/bookstore_users_api/datasources/mysql/users_db"
	"github.com/r-zareba/bookstore_users_api/logger"
	"github.com/r-zareba/bookstore_utils-go/rest_errors"
	"github.com/r-zareba/bookstore_users_api/utils/mysql_utils"
)

const (
	insertUserQuery             = "INSERT INTO users(first_name, last_name, email, date_created, status, password) VALUES(?, ?, ?, ?, ?, ?);"
	getUserQuery                = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE id=?;"
	updateQuery                 = "UPDATE users SET first_name=?, last_name=?, email=?, status=?, password=? WHERE id=?;"
	deleteQuery                 = "DELETE FROM users WHERE id=?;"
	findByStatusQuery           = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE status=?"
	findByEmailAndPasswordQuery = "SELECT id, first_name, last_name, email, date_created, status FROM users WHERE email=? AND password=?"
)

func (user *User) GetFromDB() *rest_errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(getUserQuery)
	if err != nil {
		logger.Error("Error when trying to prepare getUserQuery", err)
		return rest_errors.InternalServerError("Database error", err)
	}
	defer statement.Close()

	result := statement.QueryRow(user.Id)
	// Scan - automatically fill the object attributes
	getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if getErr != nil {
		return mysql_utils.GetMySQLError(getErr)
	}
	return nil
}

func (user *User) DeleteFromDB() *rest_errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(deleteQuery)
	if err != nil {
		logger.Error("Error when trying to prepare deleteUserQuery", err)
		return rest_errors.InternalServerError("Database error", err)
	}
	defer statement.Close()

	_, deleteErr := statement.Exec(user.Id)
	if deleteErr != nil {
		return mysql_utils.GetMySQLError(deleteErr)
	}
	return nil
}

func (user *User) FindByStatusInDB(status string) ([]User, *rest_errors.RestError) {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(findByStatusQuery)
	if err != nil {
		logger.Error("Error when trying to prepare deleteUserQuery", err)
		return nil, rest_errors.InternalServerError("Database error", err)
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
		return nil, rest_errors.NotFoundError(fmt.Sprintf("Users with status: %s not found", status))
	}
	return results, nil
}

func (user *User) SaveToDB() *rest_errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(insertUserQuery)
	if err != nil {
		logger.Error("Error when trying to prepare insertUserQuery", err)
		return rest_errors.InternalServerError("Database error", err)
	}
	defer statement.Close()

	insertResult, saveErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.DateCreated,
		user.Status, user.Password)

	if saveErr != nil {
		return mysql_utils.GetMySQLError(saveErr)
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return rest_errors.InternalServerError("Cannot get last inserted ID", err)
	}

	user.Id = userId
	return nil
}

func (user *User) UpdateInDB() *rest_errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(updateQuery)
	if err != nil {
		logger.Error("Error when trying to prepare updateQuery", err)
		return rest_errors.InternalServerError("Database error", err)
	}
	defer statement.Close()

	_, updateErr := statement.Exec(user.FirstName, user.LastName, user.Email, user.Status, user.Password, user.Id)
	if updateErr != nil {
		return mysql_utils.GetMySQLError(updateErr)
	}
	return nil
}

func (user *User) FindByEmailAndPasswordInDB() *rest_errors.RestError {
	// Prepare statement
	statement, err := users_db.ClientDB.Prepare(findByEmailAndPasswordQuery)
	if err != nil {
		logger.Error("Error when trying to prepare findByEmailAndPasswordQuery", err)
		return rest_errors.InternalServerError("Database error", err)
	}
	defer statement.Close()

	result := statement.QueryRow(user.Email, user.Password)
	// Scan - automatically fill the object attributes
	getErr := result.Scan(&user.Id, &user.FirstName, &user.LastName, &user.Email, &user.DateCreated, &user.Status)
	if getErr != nil {
		return mysql_utils.GetMySQLError(getErr)
	}
	return nil
}