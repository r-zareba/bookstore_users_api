package users_controller

import (
	"github.com/gin-gonic/gin"
	"github.com/r-zareba/bookstore_users_api/domain/users"
	"github.com/r-zareba/bookstore_users_api/services"
	"github.com/r-zareba/bookstore_users_api/utils/errors"
	"net/http"
	"strconv"
)

func Create(ctx *gin.Context) {
	var user users.User
	// Bind json directly to fill User struct fields
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid JSON body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		ctx.JSON(saveErr.Status, saveErr)
		return
	}
	ctx.JSON(http.StatusCreated, result.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func Delete(ctx *gin.Context) {
	userId, idError := parseUserId(ctx.Param("user_id"))
	if idError != nil {
		ctx.JSON(idError.Status, idError)
		return
	}

	_, deleteErr := services.DeleteUser(userId)
	if deleteErr != nil {
		ctx.JSON(deleteErr.Status, deleteErr)
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{"status": "deleted"})
}

func Get(ctx *gin.Context) {
	userId, idError := parseUserId(ctx.Param("user_id"))
	if idError != nil {
		ctx.JSON(idError.Status, idError)
		return
	}

	user, getError := services.GetUser(userId)
	if getError != nil {
		ctx.JSON(getError.Status, getError)
		return
	}

	ctx.JSON(http.StatusOK, user.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func Login(ctx *gin.Context) {
	var loginRequest users.LoginRequest
	err := ctx.ShouldBindJSON(&loginRequest)
	if err != nil {
		restErr := errors.BadRequestError("Invalid JSON body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	user, getError := services.LoginUser(loginRequest)
	if getError != nil {
		ctx.JSON(getError.Status, getError)
		return
	}
	ctx.JSON(http.StatusOK, user.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func Update(ctx *gin.Context) {
	userId, idError := parseUserId(ctx.Param("user_id"))
	if idError != nil {
		ctx.JSON(idError.Status, idError)
		return
	}

	user := users.User{Id: userId}
	// Bind json directly to fill User struct fields
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("Invalid JSON body")
		ctx.JSON(restErr.Status, restErr)
		return
	}

	isPartialUpdate := ctx.Request.Method == http.MethodPatch
	result, updateErr := services.UpdateUser(isPartialUpdate, user)
	if updateErr != nil {
		ctx.JSON(updateErr.Status, updateErr)
		return
	}
	ctx.JSON(http.StatusOK, result.Marshall(ctx.GetHeader("X-Public") == "true"))

}

func Search(ctx *gin.Context) {
	status := ctx.Query("status")

	users, err := services.Search(status)
	if err != nil {
		ctx.JSON(err.Status, err)
		return
	}
	ctx.JSON(http.StatusOK, users.Marshall(ctx.GetHeader("X-Public") == "true"))
}

func parseUserId(userIdParam string) (int64, *errors.RestError) {
	userId, idError := strconv.ParseInt(userIdParam, 10, 64)
	if idError != nil {
		return 0, errors.BadRequestError("Cannot parse user id from URL")
	}
	return userId, nil
}
