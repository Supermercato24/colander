package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermercato24/configuration"
)

func TestConst(t *testing.T) {
	assert.NotEmpty(t, configuration.DirBinStorage)
}
