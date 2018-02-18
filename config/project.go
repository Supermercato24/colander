// Package config implements configuration variables
//
// project configuration
package config

import (
	"path/filepath"

	"github.com/supermercato24/configuration"
)

const (
	// Name of the project
	Name = "colander"

	// Version of the project
	Version = "0.2.0"
)

func init() {
	configuration.DirProject = filepath.Join(configuration.DirProject, Name)
	configuration.BuildProject()
	configuration.DirBinStorage = filepath.Join(configuration.DirProject, configuration.PathStorage)
}
