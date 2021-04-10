package main

import (
	"log"

	"github.com/nicolaspearson/go.api/cmd/api/app"
)

// @title User API
// @version 1.0
// @description Swagger API documentation for the User API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email nic.s.pearson@gmail.com

// @license.name MIT
// @license.url https://github.com/nicolaspearson/go.api/blob/master/LICENSE

// @host localhost:3000
// @BasePath /
func main() {
	application := app.New()
	err := application.Build().Run()
	if err != nil {
		log.Fatalf("Application failed to start: %v", err)
	}
}
