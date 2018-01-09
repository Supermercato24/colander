// Package config implements configuration variables
//
// project configuration
package config

import (
	"path/filepath"

	"github.com/supermercato24/configuration"
)

const (
	Name    = "colander" // project name
	Version = "0.2.0"    // project version
)

func init() {
	configuration.DirProject = filepath.Join(configuration.DirProject, Name)
	configuration.BuildProject()
	configuration.DirBinStorage = filepath.Join(configuration.DirProject, configuration.PathStorage)
}
