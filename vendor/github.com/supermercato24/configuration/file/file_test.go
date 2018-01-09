package file

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	fileNameIncorrect = "foo"
	fileNameCorrect   = "file_test.go"
	dirNameIncorrect  = "bar"
	dirNameCorrect    = "file"
)

var (
	dirPathIncorrect = filepath.Join("..", dirNameIncorrect)
	dirPathCorrect   = filepath.Join("..", dirNameCorrect)
)

func TestExist(t *testing.T) {
	var fileInfo os.FileInfo
	var err error

	fileInfo, err = os.Stat(fileNameIncorrect)
	assert.Error(t, err)
	assert.Empty(t, fileInfo)
	assert.True(t, os.IsNotExist(err))

	fileInfo, err = os.Stat(fileNameCorrect)
	assert.NoError(t, err)
	assert.NotEmpty(t, fileInfo)
	assert.False(t, os.IsNotExist(err))

	fileInfo, err = Exist(dirPathIncorrect)
	assert.Error(t, err)
	assert.Empty(t, fileInfo)
	assert.True(t, os.IsNotExist(err))

	fileInfo, err = Exist(dirPathCorrect)
	assert.NoError(t, err)
	assert.NotEmpty(t, fileInfo)
	assert.False(t, os.IsNotExist(err))
}

func TestIsFileIncorrect(t *testing.T) {
	fileInfo, err := os.Stat(fileNameIncorrect)
	assert.Error(t, err)
	assert.Empty(t, fileInfo)
	assert.True(t, os.IsNotExist(err))

	exist, fileInfo, err := IsFile(fileNameIncorrect)
	assert.False(t, exist)
	assert.Empty(t, fileInfo)
	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err))
}

func TestIsFile(t *testing.T) {
	fileInfo, err := os.Stat(fileNameCorrect)
	assert.NoError(t, err)
	assert.NotEmpty(t, fileInfo)
	assert.False(t, os.IsNotExist(err))

	mode := fileInfo.Mode()
	assert.True(t, mode.IsRegular())
	assert.False(t, mode.IsDir())

	exist, fileInfo, err := IsFile(fileNameCorrect)
	assert.NotEmpty(t, fileInfo)
	assert.True(t, exist)
	assert.NoError(t, err)

	exist, fileInfo, err = IsDirectory(fileNameCorrect)
	assert.NotEmpty(t, fileInfo)
	assert.False(t, exist)
	assert.NoError(t, err)
}

func TestIsDirectoryIncorrect(t *testing.T) {
	fileInfo, err := os.Stat(dirNameIncorrect)
	assert.Error(t, err)
	assert.Empty(t, fileInfo)
	assert.True(t, os.IsNotExist(err))

	exist, fileInfo, err := IsDirectory(dirNameIncorrect)
	assert.False(t, exist)
	assert.Empty(t, fileInfo)
	assert.Error(t, err)
	assert.True(t, os.IsNotExist(err))
}

func TestIsDirectory(t *testing.T) {
	fileInfo, err := os.Stat(dirPathCorrect)
	assert.NoError(t, err)
	assert.NotEmpty(t, fileInfo)
	assert.False(t, os.IsNotExist(err))

	mode := fileInfo.Mode()
	assert.False(t, mode.IsRegular())
	assert.True(t, mode.IsDir())

	exist, fileInfo, err := IsFile(dirPathCorrect)
	assert.NotEmpty(t, fileInfo)
	assert.False(t, exist)
	assert.NoError(t, err)

	exist, fileInfo, err = IsFile(dirPathCorrect)
	assert.NotEmpty(t, fileInfo)
	assert.False(t, exist)
	assert.NoError(t, err)
}

func TestGetFileDescriptorIncorrect(t *testing.T) {
	fd, fileInfo, err := GetFileDescriptor(fileNameIncorrect)
	assert.Empty(t, fileInfo)
	assert.Empty(t, fd)
	assert.Error(t, err)
}

func TestGetFileDescriptor(t *testing.T) {
	fd, fileInfo, err := GetFileDescriptor(fileNameCorrect)
	assert.NotEmpty(t, fileInfo)
	assert.NotEmpty(t, fd)
	assert.NoError(t, err)

	assert.NoError(t, fd.Close())
}
