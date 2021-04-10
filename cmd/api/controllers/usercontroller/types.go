package usercontroller

import (
	"github.com/gin-gonic/gin"
	"github.com/nicolaspearson/go.api/cmd/api/internal/application/userservice"
)

type IUserController interface {
	Init(e *gin.Engine)
}

type UserController struct {
	service userservice.IUserService
}
