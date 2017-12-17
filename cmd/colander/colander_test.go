package main

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"supermercato24.it/colander/config"
	"supermercato24.it/configuration"
)

func TestColanderWithScreen(t *testing.T) {
	assert.Exactly(t, config.Name, config.Name, "load init")

	Colander(&ColanderOptions{
		dir:  filepath.Join(configuration.DirBinStorage, configuration.PathLogs),
		show: true,
	})
}

func TestColanderWithoutScreen(t *testing.T) {
	assert.Exactly(t, config.Name, config.Name, "load init")

	Colander(&ColanderOptions{
		dir:  filepath.Join(configuration.DirBinStorage, configuration.PathLogs),
		show: false,
	})
}
