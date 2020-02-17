package users

import (
	"github.com/gin-gonic/gin"
	"github.com/r-zareba/bookstore_users_api/domain/users"
	"github.com/r-zareba/bookstore_users_api/services"
	"github.com/r-zareba/bookstore_users_api/utils/errors"
	"net/http"
	"strconv"
)

func CreateUser(ctx *gin.Context) {
	var user users.User
	// Bind json directly to fill User struct fields
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid JSON body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	//fmt.Println(user)
	//bytes, err := ioutil.ReadAll(ctx.Request.Body)
	//if err != nil {
	//
	//	return
	//}
	//
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println(err.Error())
	//
	//	return
	//}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result)
}

func GetUser(ctx *gin.Context) {
	userId, idError := strconv.ParseInt(ctx.Param("user_id"), 10, 64)
	if idError != nil {
		err := errors.BadRequestError("Cannot parse user if from URL")
		ctx.JSON(err.Status, err)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil {
		ctx.JSON(getError.Status, getError)
		return
	}

	ctx.JSON(http.StatusOK, user)

}
