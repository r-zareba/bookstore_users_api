package mysql_utils

import (
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/r-zareba/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	noRowsError = "no rows in result set"
)

func GetMySQLError(err error) *rest_errors.RestError {
	// Cast to MySQL Error Type for easier handling
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), noRowsError) {
			return rest_errors.NotFoundError("No record matching given id")
		}
		return rest_errors.InternalServerError("Error parsing database response", errors.New("Database error"))
	}
	return rest_errors.InternalServerError(fmt.Sprintf("Internal Database Error: %s", sqlErr.Error()),
		errors.New("Database error"))
}
