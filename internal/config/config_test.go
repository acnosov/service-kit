package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	cfg, err := Load()
	assert.Nil(t, err)
	assert.Equal(t, "service-kit", cfg.Service.Name)
	assert.Equal(t, "dev", cfg.Service.Environment)
}
