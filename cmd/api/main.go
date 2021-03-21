package main

import (
	"log"

	"github.com/nicolaspearson/go.api/cmd/api/config"
)

func main() {
	if err := config.LoadConfig("./config"); err != nil {
		log.Fatalf("Invalid application configuration: %v", err)
	}
	log.Printf("Environment: %s", config.Vars.Environment)
	log.Printf("ReleaseVersion: %s", config.Vars.ReleaseVersion)
	log.Printf("Version: %s", config.Vars.Version)
}
