package file

import (
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/supermercato24/colander/config"
	"github.com/supermercato24/configuration"
)

const (
	dailyLog0   = "day"
	dailyLog1   = "d2"
	monthlyLog0 = "month"
)

func globLogs(t *testing.T) {
	var day, d2, month bool

	Glob(filepath.Join(configuration.DirBinStorage, configuration.PathLogs), "", func(matches *GlobMatches) {
		for filesKey, match := range matches.Files {
			if filesKey == dailyLog0 {
				assert.Exactly(t, dailyLog0, filesKey)
				day = true
			} else if filesKey == dailyLog1 {
				assert.Exactly(t, dailyLog1, filesKey)
				d2 = true
			} else {
				assert.Exactly(t, monthlyLog0, filesKey)
				month = true
			}
			assert.Exactly(t, match.Category, filesKey)

			for logsKey, match := range match.Logs {
				assert.Exactly(t, int64(2017), match.Year)
				if filesKey == monthlyLog0 {
					assert.Exactly(t, int64(0), match.Day)
				} else {
					assert.Exactly(t, int64(12), match.Month)
				}
				assert.IsType(t, time.Time{}, logsKey)
				assert.NotZero(t, len(match.Paths))

				for _, match := range match.Paths {
					assert.True(t, filepath.IsAbs(match))
				}
			}
		}
	})

	assert.True(t, day)
	assert.True(t, d2)
	assert.True(t, month)
}

func TestGlob2(t *testing.T) {
	assert.Exactly(t, config.Name, config.Name, "load init")

	assert.True(t, t.Run("setUp0", func(t *testing.T) { logSetUp(t, 0) }))
	assert.True(t, t.Run("setUp3", func(t *testing.T) { logSetUp(t, 3) }))
	assert.True(t, t.Run("setUp4", func(t *testing.T) { logSetUp(t, 4) }))

	assert.True(t, t.Run("globLogs", globLogs))
}
