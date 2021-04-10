package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nicolaspearson/go.api/cmd/api/internal/application/userservice"
)

func New(service userservice.IUserService) IUserController {
	return &UserController{service: service}
}

func (controller UserController) Init(e *gin.Engine) {
	group := e.Group("api/users")
	{
		group.POST("", controller.createHandler)
		group.GET("", controller.getAllHandler)
		group.GET("/:uuid", controller.getByUuidHandler)
	}
}

// @Summary Create a user
// @Description Saves a new user to the database
// @Accept  json
// @Produce  json
// @tags UserController
// @param UserCreateRequestDto body userservice.UserCreateRequestDto true "Create a user"
// @Success 201 {string} string	"Success"
// @Router /api/users [post]
func (controller UserController) createHandler(c *gin.Context) {
	var userDto userservice.UserCreateRequestDto
	if err := c.ShouldBindJSON(&userDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.service.Create(userDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "An error occurred while creating the user!"})
		return
	}

	c.Status(http.StatusCreated)
}

// @Summary Get all users
// @Description Retrieves all users from the database
// @Accept  json
// @Produce  json
// @tags UserController
// @Success 200 {object} []userservice.UserDto "Success"
// @Router /api/users [get]
func (controller UserController) getAllHandler(c *gin.Context) {
	users, err := controller.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred retrieving users!"})
		return
	}

	c.JSON(http.StatusOK, users)
}

// @Summary Get user by uuid
// @Description Retrieves the user identified by the provided UUID from the database
// @Accept  json
// @Produce  json
// @tags UserController
// @Success 200 {object} userservice.UserDto "Success"
// @Router /api/users/{uuid} [get]
func (controller UserController) getByUuidHandler(c *gin.Context) {
	uuid, parseError := uuid.Parse(c.Param("uuid"))
	if parseError != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid uuid provided!"})
		return
	}

	user, err := controller.service.GetByUuid(uuid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred retrieving the requested user!"})
		return
	}

	c.JSON(http.StatusOK, user)
}
