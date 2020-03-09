package services

import (
	"github.com/r-zareba/bookstore_users_api/domain/users"
	"github.com/r-zareba/bookstore_users_api/utils/crypto_utils"
	"github.com/r-zareba/bookstore_users_api/utils/date_utils"
	"github.com/r-zareba/bookstore_users_api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.DateCreated = date_utils.GetNowTime()
	user.Password = crypto_utils.GetMd5(user.Password)

	err := user.SaveToDB()
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestError) {
	user := users.User{Id: userId}
	if err := user.GetFromDB(); err != nil {
		return nil, err
	}
	return &user, nil
}

func LoginUser(request users.LoginRequest) (*users.User, *errors.RestError) {
	user := users.User{
		Email: request.Email,
		Password: crypto_utils.GetMd5(request.Password)}
	if err := user.FindByEmailAndPasswordInDB(); err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUser(userId int64) (*users.User, *errors.RestError) {
	currentUser, getErr := GetUser(userId)
	if getErr != nil {
		return nil, getErr
	}

	updateErr := currentUser.DeleteFromDB()
	if updateErr != nil {
		return nil, updateErr
	}
	return currentUser, nil
}

func UpdateUser(isPartial bool, user users.User) (*users.User, *errors.RestError) {
	currentUser, getErr := GetUser(user.Id)
	if getErr != nil {
		return nil, getErr
	}

	if isPartial {
		currentUser.PartialUpdateFields(user)
	} else {
		if err := user.Validate(); err != nil {
			return nil, err
		}
		currentUser.UpdateFields(user)
	}

	updateErr := currentUser.UpdateInDB()
	if updateErr != nil {
		return nil, updateErr
	}
	return currentUser, nil
}

func Search(status string) (users.Users, *errors.RestError) {
	var dao users.User
	return dao.FindByStatusInDB(status)
}
