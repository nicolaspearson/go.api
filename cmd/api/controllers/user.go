package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/nicolaspearson/go.api/cmd/api/db/daos"
	"github.com/nicolaspearson/go.api/cmd/api/services"
)

// GetUser godoc
// @Summary Retrieves the user identified by the provided ID
// @Produce json
// @Param id path integer true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
// @Security ApiKeyAuth
func GetById(c *gin.Context) {
	s := services.NewUserService(daos.NewUserDAO())
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if user, err := s.GetById(uint(id)); err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, user)
	}
}
