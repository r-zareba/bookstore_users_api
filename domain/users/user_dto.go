package users

import (
	"errors"
	"github.com/r-zareba/bookstore_utils-go/rest_errors"
	"strings"
)

const (
	ActiveUserStatus = "active"
	BannedUserStatus = "banned"
)

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"` // Don't bind from/to json
}

type Users []User

func (user *User) Validate() *rest_errors.RestError {
	// Check user fields
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)
	user.Status = ActiveUserStatus

	emailErr := user.checkEmail()
	if emailErr != nil {
		return emailErr
	}

	passwordErr := user.checkPassword()
	if passwordErr != nil {
		return passwordErr
	}

	return nil
}

func (user *User) PartialUpdateFields(other User) {
	if other.FirstName != "" {
		user.FirstName = other.FirstName
	}
	if other.LastName != "" {
		user.LastName = other.LastName
	}
	if other.Email != "" {
		user.Email = other.Email
	}
	if other.Password != "" {
		user.Password = other.Password
	}
	if other.Status != "" {
		user.Status = other.Status
	}
}

func (user *User) UpdateFields(other User) {
	user.FirstName = other.FirstName
	user.LastName = other.LastName
	user.Email = other.Email
}

func (user *User) checkEmail() *rest_errors.RestError {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.BadRequestError("Invalid email address")
	}
	return nil
}

func (user *User) checkPassword() *rest_errors.RestError {
	user.Password = strings.TrimSpace(user.Password)
	if len(user.Password) < 5 {
		return rest_errors.InternalServerError("Password too short!", errors.New("Password too short"))
	}
	return nil
}

