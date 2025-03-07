package controller

import (
	"fmt"

	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/deadman360/crud_portifolio/src/controller/routes/model/request"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var UserRequest request.UserRequest
	if err := c.ShouldBindBodyWithJSON(&UserRequest); err != nil {
		restErr := rest_err.NewBadRequestError(
			fmt.Sprintf("Campos Incorretos", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}
	fmt.Println(UserRequest)
}
