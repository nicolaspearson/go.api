package indexcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func New() IIndexController {
	return &IndexController{}
}

func (controller IndexController) Init(e *gin.Engine) {
	e.GET("/", controller.indexHandler)
}

// @Summary Redirects to the Swagger UI
// @Description Redirects to the Swagger UI
// @Accept  json
// @Produce  json
// @tags IndexController
// @Success 308 {string} string	"Redirect"
// @Router / [get]
func (controller IndexController) indexHandler(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "/swagger/index.html")
}
