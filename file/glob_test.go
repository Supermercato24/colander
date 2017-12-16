package file

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"supermercato24.it/colander/config"
	"supermercato24.it/configuration"
)

func TestGlob(t *testing.T) {
	assert.Exactly(t, config.Name, config.Name, "load init")

	Glob(filepath.Join(configuration.DirBinStorage, configuration.PathLogs), func(matches GlobMatches) {
		for filesKey, match := range matches.Files {
			if filesKey == "day" {
				assert.Exactly(t, "day", filesKey)
			} else if filesKey == "d2" {
				assert.Exactly(t, "d2", filesKey)
			} else {
				assert.Exactly(t, "month", filesKey)
			}
			assert.Exactly(t, match.Category, filesKey)

			for logsKey, match := range match.Logs {
				assert.Exactly(t, int64(2017), match.Year)
				assert.IsType(t, time.Time{}, logsKey)
				assert.NotZero(t, len(match.Logs))

				for _, match := range match.Logs {
					assert.True(t, filepath.IsAbs(match))
				}
			}
		}
	})
}
