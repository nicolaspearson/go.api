package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	services "github.com/nicolaspearson/go.api/cmd/api/services"
)

type UserController struct {
	userService services.IUserService
}

type IUserController interface {
	GetById(c *gin.Context)
}

func NewUserController(s services.IUserService) IUserController {
	return &UserController{
		userService: s,
	}
}

// GetUser godoc
// @Summary Retrieves the user identified by the provided ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} db.User
// @Router /users/{id} [get]
func (c *UserController) GetById(context *gin.Context) {
	id, _ := strconv.ParseUint(context.Param("id"), 10, 32)
	if user, err := c.userService.GetById(uint(id)); err != nil {
		context.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		context.JSON(http.StatusOK, user)
	}
}
