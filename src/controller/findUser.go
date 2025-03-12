package controller

import (
	"github.com/deadman360/crud_portifolio/src/configuration/rest_err"
	"github.com/gin-gonic/gin"
)

func (uc *userControllerInterface) FindUserById(c *gin.Context) {
	err := rest_err.NewBadRequestError("ID Not Found")
	c.JSON(err.Code, err)
}
func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {

}
