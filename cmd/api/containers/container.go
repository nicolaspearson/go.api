package containers

import (
	"fmt"
	"strings"

	"github.com/ory/dockertest"
	"github.com/ory/dockertest/docker"
)

func IsRunning(pool dockertest.Pool, imagename string) bool {
	dockerContainers, _ := pool.Client.ListContainers(docker.ListContainersOptions{
		All: false,
	})

	for _, dockerContainer := range dockerContainers {
		for _, name := range dockerContainer.Names {
			if strings.Contains(name, imagename) {
				fmt.Printf("%s image is running..\n", dockerContainer.Image)
				return true
			}
		}
	}

	return false
}
