package configuration

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfigVariables(t *testing.T) {
	assert.True(t, filepath.IsAbs(DirProject))

	assert.Contains(t, DirScript, DirProject)
	assert.Contains(t, DirScript, PathScript)
	assert.NotEqual(t, DirScript, DirProject)
	assert.Exactly(t, filepath.Dir(DirScript), DirProject)
	assert.True(t, filepath.IsAbs(DirScript))

	assert.NotContains(t, DirProjectApp, DirScript, "python in python")
	assert.Contains(t, DirProjectApp, DirProject, "php in php")
	assert.Contains(t, DirProjectApp, PathApp)
	assert.NotEqual(t, DirProjectApp, DirProject)
	assert.NotEqual(t, DirProjectApp, DirScriptApp)
	assert.Exactly(t, filepath.Dir(DirProjectApp), DirProject)
	assert.True(t, filepath.IsAbs(DirProjectApp))

	assert.Contains(t, DirScriptApp, DirProject, "python in python")
	assert.Contains(t, DirScriptApp, DirScript, "php in php")
	assert.Contains(t, DirScriptApp, PathApp)
	assert.NotEqual(t, DirScriptApp, DirProject)
	assert.NotEqual(t, DirScriptApp, DirProjectApp)
	assert.Exactly(t, filepath.Dir(DirScriptApp), DirScript)
	assert.True(t, filepath.IsAbs(DirScriptApp))

	assert.NotContains(t, DirProjectConfig, DirScript, "python in python")
	assert.Contains(t, DirProjectConfig, DirProject, "php in php")
	assert.NotContains(t, DirProjectConfig, PathApp)
	assert.Contains(t, DirProjectConfig, PathConfig)
	assert.NotEqual(t, DirProjectConfig, DirProject)
	assert.NotEqual(t, DirProjectConfig, DirScriptConfig)
	assert.Exactly(t, filepath.Dir(DirProjectConfig), DirProject)
	assert.True(t, filepath.IsAbs(DirProjectConfig))

	assert.Contains(t, DirScriptConfig, DirProject, "python in python")
	assert.Contains(t, DirScriptConfig, DirScript, "php in php")
	assert.Contains(t, DirScriptConfig, PathApp)
	assert.Contains(t, DirProjectConfig, PathConfig)
	assert.NotEqual(t, DirScriptConfig, DirProject)
	assert.NotEqual(t, DirScriptConfig, DirProjectConfig)
	assert.Exactly(t, filepath.Dir(DirScriptConfig), DirScriptApp)
	assert.True(t, filepath.IsAbs(DirScriptConfig))

	assert.NotContains(t, DirProjectStorage, DirScript, "python in python")
	assert.Contains(t, DirProjectStorage, DirProject, "php in php")
	assert.Contains(t, DirProjectStorage, PathStorage)
	assert.NotEqual(t, DirProjectStorage, DirProject)
	assert.NotEqual(t, DirProjectStorage, DirScriptStorage)
	assert.Exactly(t, filepath.Dir(DirProjectStorage), DirProject)
	assert.True(t, filepath.IsAbs(DirProjectStorage))

	assert.Contains(t, DirScriptStorage, DirProject, "python in python")
	assert.Contains(t, DirScriptStorage, DirScript, "php in php")
	assert.Contains(t, DirScriptStorage, PathStorage)
	assert.NotEqual(t, DirScriptStorage, DirProject)
	assert.NotEqual(t, DirScriptStorage, DirProjectStorage)
	assert.Exactly(t, filepath.Dir(DirScriptStorage), DirScript)
	assert.True(t, filepath.IsAbs(DirScriptStorage))

	assert.True(t, filepath.IsAbs(DirBinStorage))
}

func TestConfigVariablesRebuild(t *testing.T) {
	BuildProject()
	TestConfigVariables(t)
}

func TestConfigCfgWithMissingFile(t *testing.T) {
	fileName := "foo"

	fileInfo, err := os.Stat(fileName)
	assert.Empty(t, fileInfo)
	assert.NotEmpty(t, fileName)
	assert.Error(t, err)

	//assert.Panics(t, func() { cfg(fileName, "bar") })
}

func TestConfigCfgWithWrongRegexp(t *testing.T) {
	testPath := filepath.Join(DirCwd, "config_test.go")

	fileInfo, err := os.Stat(testPath)

	assert.NotEmpty(t, fileInfo)
	assert.NotEmpty(t, testPath)
	assert.NoError(t, err)

	mode := fileInfo.Mode()
	assert.True(t, mode.IsRegular())
	assert.False(t, mode.IsDir())

	assert.True(t, assert.NotPanics(t, func() { cfg(testPath, "bar") }))

	object, err := cfg(testPath, "bar")
	assert.Empty(t, object)
	assert.NoError(t, err)

	assert.True(t, assert.NotPanics(t, func() { cfg(testPath, `(?P<foobar>[\d]+)`) }))
}
