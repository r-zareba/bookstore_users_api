package users

import (
	"github.com/gin-gonic/gin"
	"github.com/r-zareba/bookstore_users_api/domain/users"
	"github.com/r-zareba/bookstore_users_api/errors"
	"github.com/r-zareba/bookstore_users_api/services"
	"net/http"
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
	//	// TODO Handle error
	//	return
	//}
	//
	//if err := json.Unmarshal(bytes, &user); err != nil {
	//	fmt.Println(err.Error())
	//	// TODO Handle error
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

}
