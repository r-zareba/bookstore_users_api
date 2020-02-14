package services

import (
	"github.com/r-zareba/bookstore_users_api/domain/users"
	"github.com/r-zareba/bookstore_users_api/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	return &user, nil
}
