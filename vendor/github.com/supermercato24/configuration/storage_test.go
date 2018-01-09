package configuration

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	storagePath   = filepath.Join(PathScript, PathApp, PathConfig, storageFile)
	storageMockup = []byte(`ENV = {
    'foobar': {
        'token': '8c-Dp'
    }
}`)
)

func StorageModel(t *testing.T, storage *storage) {
	// storage
	assert.NotEmpty(t, storage.Token())
}

func StorageMockup(t *testing.T) {
	cwd, _ := os.Getwd()
	DirProject = cwd
	BuildProject()

	fd, err := os.Create(storagePath)
	assert.NoError(t, err)
	fd.Write(storageMockup)
	assert.NoError(t, fd.Close())
}

func TestStorage(t *testing.T) {
	assert.True(t, t.Run("mockup", StorageMockup))
	assert.True(t, assert.NotPanics(t, func() { StorageConfiguration() }))
	assert.NotPanics(t, func() { StorageModel(t, StorageConfiguration()) })
	assert.NoError(t, os.Remove(storagePath))
}

func BenchmarkStorage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StorageConfiguration()
	}
}
