package main

import (
	"testing"

	"github.com/nicolaspearson/go.api/cmd/api/config"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	if err := config.LoadConfig("../../config"); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "development", config.Vars.Environment)
	assert.Equal(t, "unknown", config.Vars.ReleaseVersion)
	assert.Equal(t, "unknown", config.Vars.Version)
}
