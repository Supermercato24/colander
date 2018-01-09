package file

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/supermercato24/colander/config"
	"github.com/supermercato24/configuration"
)

const (
	numberOfLines   = 18
	dailyLogFile0   = "d2-2017-12-11.log"
	dailyLogFile1   = "d2-2017-12-11-ws01.log"
	dailyLogFile2   = "d2-2017-12-11-ws02.log"
	dailyLogFile3   = "day-2017-12-11.log"
	monthlyLogFile0 = "month-2017-12.log"
)

var (
	dailyLogBody0 = []byte(`[2017-12-01 23:16:20] d2-201712-11
[2017-12-01 23:21:00] d2-201712-11
[2017-12-01 23:23:40] d2-201712-11
[2017-12-01 23:23:50] d2-201712-11
[2017-12-01 23:24:20] d2-201712-11
`)
	dailyLogBody1 = []byte(`[2017-12-01 23:16:21] d2-201712-10-ws01
[2017-12-01 23:21:01] d2-201712-11-ws01
[2017-12-01 23:23:41] d2-201712-11-ws01
[2017-12-01 23:23:51] d2-201712-11-ws01
[2017-12-01 23:24:21] d2-201712-11-ws01
[2017-12-01 23:25:21] d2-201712-10-ws01
`)
	dailyLogBody2 = []byte(`[2017-12-01 23:16:22] d2-201712-10-ws02
[2017-12-01 23:21:02] d2-201712-11-ws02
[2017-12-01 23:23:42] d2-201712-11-ws02
[2017-12-01 23:23:52] d2-201712-11-ws02
[2017-12-01 23:24:22] d2-201712-11-ws02
[2017-12-01 23:25:22] d2-201712-10-ws02
[2017-12-01 23:26:22] d2-201712-10-ws01
`)
)

func TestDetectTimestamp(t *testing.T) {
	var validTimestamp []byte
	var expectedTimestamp time.Time
	var result time.Time

	validTimestamp = []byte("[2017-12-01 23:50:07]")
	expectedTimestamp = time.Date(2017, time.December, 1, 23, 50, 07, 0, time.UTC)
	result = detectTimestamp(validTimestamp)
	assert.Exactly(t, expectedTimestamp, result)

	validTimestamp = []byte("foo2017-12-01 23:50:07bar")
	expectedTimestamp = time.Date(2017, time.December, 1, 23, 50, 07, 0, time.UTC)
	result = detectTimestamp(validTimestamp)
	assert.Exactly(t, expectedTimestamp, result)

	validTimestamp = []byte("foo2017-12-01bar23:50:")
	expectedTimestamp = time.Time{}
	result = detectTimestamp(validTimestamp)
	assert.Exactly(t, expectedTimestamp, result)

	validTimestamp = []byte("foo2017-12-01bar23:50: 2017-12-01 23:50:07")
	expectedTimestamp = time.Date(2017, time.December, 1, 23, 50, 07, 0, time.UTC)
	result = detectTimestamp(validTimestamp)
	assert.Exactly(t, expectedTimestamp, result)
}

func logSetUp(t *testing.T, logNumber int) {
	var dailyLogFile string
	var dailyLogBody []byte

	switch logNumber {
	case 0:
		dailyLogFile = dailyLogFile0
		dailyLogBody = dailyLogBody0
	case 1:
		dailyLogFile = dailyLogFile1
		dailyLogBody = dailyLogBody1
	case 2:
		dailyLogFile = dailyLogFile2
		dailyLogBody = dailyLogBody2
	case 3:
		dailyLogFile = dailyLogFile3
		dailyLogBody = dailyLogBody2
	case 4:
		dailyLogFile = monthlyLogFile0
		dailyLogBody = dailyLogBody2
	}
	assert.NotEmpty(t, dailyLogFile)
	assert.NotEmpty(t, dailyLogBody)

	outputFile := filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile)
	os.Remove(outputFile)
	fd, err := os.OpenFile(outputFile, os.O_CREATE|os.O_WRONLY, 0644)
	assert.NoError(t, err)

	//_, err = fd.Seek(int64(64), 0)
	//assert.NoError(t, err)

	n, err := fd.Write(dailyLogBody)
	assert.NoError(t, err)
	assert.NotZero(t, n)

	err = fd.Close()
	assert.NoError(t, err)
}

func logReadline(t *testing.T) {
	Glob(filepath.Join(configuration.DirBinStorage, configuration.PathLogs), "", func(matches *GlobMatches) {
		match := matches.Files[dailyLog1]
		assert.Exactly(t, dailyLog1, match.Category)

		for _, match := range match.Logs {
			if match.Day != 11 {
				continue
			}

			assert.Exactly(t, int64(11), match.Day)
			logs := LogReadLines(match.Paths)
			assert.NotEmpty(t, logs)

			assert.Len(t, logs, numberOfLines)

			break
		}
	})

	Glob(filepath.Join(configuration.DirBinStorage, configuration.PathLogs), "*2017-12-11*.log", func(matches *GlobMatches) {
		match := matches.Files[dailyLog1]
		assert.Exactly(t, dailyLog1, match.Category)

		for _, match := range match.Logs {
			assert.Exactly(t, int64(11), match.Day)
			logs := LogReadLines(match.Paths)
			assert.NotEmpty(t, logs)

			assert.Len(t, logs, numberOfLines)
		}
	})
}

func TestLog(t *testing.T) {
	assert.Exactly(t, config.Name, config.Name, "load init")

	assert.True(t, t.Run("setUp0", func(t *testing.T) { logSetUp(t, 0) }))
	assert.True(t, t.Run("setUp1", func(t *testing.T) { logSetUp(t, 1) }))
	assert.True(t, t.Run("setUp2", func(t *testing.T) { logSetUp(t, 2) }))

	assert.True(t, t.Run("logReadline", logReadline))
}
