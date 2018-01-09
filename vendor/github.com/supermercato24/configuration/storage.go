// Package configuration implements methods for get storage variables.
//
// script/app/config/storage file.
package configuration

import (
	"errors"
	"path/filepath"
	"sync"
)

const (
	storageFile   = "storage.py"
	storageRegexp = `(?:[\s]{2}'token':[\s]{1}'(?P<token>[\S]+)')`
)

var (
	structStorage   storage
	errEmptyStorage = errors.New("empty Storage configuration")
	onceStorage     sync.Once
)

type storage struct {
	token string
}

func (s storage) Token() string {
	return s.token
}

func storageLoad() {
	buffer, err := cfg(filepath.Join(DirScriptConfig, storageFile), storageRegexp)

	if err != nil {
		panic(errEmptyStorage)
	}

	structStorage = storage{
		token: string(buffer["token"]),
	}

	if (storage{}) == structStorage {
		panic(errEmptyStorage)
	}
}

// StorageConfiguration expose STORAGE configuration data.
func StorageConfiguration() *storage {
	onceStorage.Do(func() {
		storageLoad()
	})

	return &structStorage
}
