package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	if err := LoadConfig("../../config"); err != nil {
		assert.Error(t, err)
	}
	assert.Equal(t, "development", Vars.Environment)
	assert.Equal(t, "unknown", Vars.ReleaseVersion)
	assert.Equal(t, "unknown", Vars.Version)
}
