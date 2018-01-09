// Package file implements methods to check if file is valid.
package file

import (
	"errors"
	"os"
)

var (
	errLocalFile = errors.New("localPath is not a file")
)

// Exist check if path exist.
func Exist(filePath string) (fileInfo os.FileInfo, err error) {
	fileInfo, err = os.Stat(filePath)
	return
}

// IsFile check if path exist and is a file.
func IsFile(filePath string) (exist bool, fileInfo os.FileInfo, err error) {
	fileInfo, err = Exist(filePath)

	if err == nil {
		if mode := fileInfo.Mode(); mode.IsRegular() {
			exist = true
		}
	}

	return
}

// IsDirectory check if path exist and is a directory.
func IsDirectory(filePath string) (exist bool, fileInfo os.FileInfo, err error) {
	fileInfo, err = Exist(filePath)

	if err == nil {
		if mode := fileInfo.Mode(); mode.IsDir() {
			exist = true
		}
	}

	return
}

// GetFileDescriptor of a file without closing it.
func GetFileDescriptor(localPath string) (fd *os.File, fileInfo os.FileInfo, err error) {
	// fileInfo check is file and exists
	exist, fileInfo, err := IsFile(localPath)

	if err != nil {
		// pass
	} else if !exist {
		err = errLocalFile
	} else {
		fd, err = os.Open(localPath)
	}

	return
}
