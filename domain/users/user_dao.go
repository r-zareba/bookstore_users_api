// The only place to execute Database commands

package users

import (
	"fmt"
	"github.com/r-zareba/bookstore_users_api/errors"
)

var usersDB = make(map[int64]*User)

// TODO Refactor? GetFromDB, FillFromDB
func (user *User) Get() *errors.RestError {
	userId := user.Id
	foundUser := usersDB[userId]
	if foundUser == nil {
		return errors.NotFoundError(fmt.Sprintf("User of id: %d not found", userId))
	}

	user.Id = foundUser.Id
	user.FirstName = foundUser.FirstName
	user.LastName = foundUser.LastName
	user.Email = foundUser.Email
	user.DateCreated = foundUser.DateCreated
	return nil
}

func (user *User) Save() *errors.RestError {
	fmt.Println(usersDB)
	currentUser := usersDB[user.Id]
	if currentUser != nil {
		if currentUser.Email == user.Email {
			return errors.BadRequestError(fmt.Sprintf("User with email: %s already registered", user.Email))
		}
		return errors.BadRequestError(fmt.Sprintf("User of ID:%d already exists in Database", user.Id))
	}
	usersDB[user.Id] = user
	return nil
}
