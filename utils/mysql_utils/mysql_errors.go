package mysql_utils

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/r-zareba/bookstore_users_api/utils/errors"
	"strings"
)

const (
	noRowsError = "no rows in result set"
)

func GetMySQLError(err error) *errors.RestError {
	// Cast to MySQL Error Type for easier handling
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return errors.NotFoundError("No record matching given id")
		}
		return errors.InternalServerError("Error parsing database response")
	}
	return errors.InternalServerError(fmt.Sprintf("Internal Database Error: %s", sqlErr.Error()))
}
