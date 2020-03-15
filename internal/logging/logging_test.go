package logging

import (
	"github.com/aibotsoft/service-kit/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}
	log := New(cfg)
	assert.NotEmpty(t, log)
}
