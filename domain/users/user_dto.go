package users

import (
	"github.com/r-zareba/bookstore_users_api/errors"
	"strings"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestError {
	// Check user fields
	if err := user.checkEmail(); err != nil {
		return err
	}
	return nil
}

func (user *User) checkEmail() *errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.BadRequestError("Invalid email address")
	}
	return nil
}
