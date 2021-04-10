package app

import "github.com/gin-gonic/gin"

func (application *Application) AddRouter() {
	// Set gin to release mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a router with logger and recovery middleware attached
	application.engine = gin.Default()
}
