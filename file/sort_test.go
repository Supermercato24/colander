package file

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"supermercato24.it/colander/config"
	"supermercato24.it/configuration"
)

func TestSortByTimestamp(t *testing.T) {
	assert.Exactly(t, config.Name, config.Name, "load init")

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
			timestampTouchpoint0 := logs[0].Timestamp
			timestampTouchpoint1 := logs[numberOfLines/2-1].Timestamp
			timestampTouchpoint2 := logs[numberOfLines/3-1].Timestamp
			timestampTouchpoint3 := logs[numberOfLines/1-1].Timestamp

			sortedLogs := SortByTimestamp(logs)

			assert.Len(t, sortedLogs, numberOfLines)
			sortedTimestampTouchpoint0 := sortedLogs[0].Timestamp
			sortedTimestampTouchpoint1 := sortedLogs[numberOfLines/2-1].Timestamp
			sortedTimestampTouchpoint2 := sortedLogs[numberOfLines/3-1].Timestamp
			sortedTimestampTouchpoint3 := sortedLogs[numberOfLines/1-1].Timestamp
			assert.True(t,
				(timestampTouchpoint0.Unix() != sortedTimestampTouchpoint0.Unix()) ||
					(timestampTouchpoint1.Unix() != sortedTimestampTouchpoint1.Unix()) ||
					(timestampTouchpoint2.Unix() != sortedTimestampTouchpoint2.Unix()) ||
					(timestampTouchpoint3.Unix() != sortedTimestampTouchpoint3.Unix()))

			assert.Exactly(t, &sortedLogs, &logs, "reference")
			assert.True(t, (&sortedTimestampTouchpoint0 != &timestampTouchpoint0) ||
				(&sortedTimestampTouchpoint1 != &timestampTouchpoint1) ||
				(&sortedTimestampTouchpoint2 != &timestampTouchpoint2) ||
				(&sortedTimestampTouchpoint3 != &timestampTouchpoint3))

			resultPath := filepath.Join(configuration.DirBinStorage, configuration.PathLogs, "test_result.log")
			err := LogWriteLines(resultPath, sortedLogs)
			assert.NoError(t, err)

			fd, err := os.Open(resultPath)
			assert.NoError(t, err)

			ops := 0
			reader := bufio.NewReader(fd)

			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					break
				}
				ops += 1
				assert.NoError(t, err)

				if ops == 1 {
					assert.Exactly(t, detectTimestamp(line), sortedTimestampTouchpoint0)
				}
			}
			assert.Exactly(t, ops, numberOfLines)
			fd.Close()

			break
		}
	})

	Glob(filepath.Join(configuration.DirBinStorage, configuration.PathLogs), "*2017-12-11*.log", func(matches *GlobMatches) {
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
			timestampTouchpoint0 := logs[0].Timestamp
			timestampTouchpoint1 := logs[numberOfLines/2-1].Timestamp
			timestampTouchpoint2 := logs[numberOfLines/3-1].Timestamp
			timestampTouchpoint3 := logs[numberOfLines/1-1].Timestamp

			sortedLogs := SortByTimestamp(logs)

			assert.Len(t, sortedLogs, numberOfLines)
			sortedTimestampTouchpoint0 := sortedLogs[0].Timestamp
			sortedTimestampTouchpoint1 := sortedLogs[numberOfLines/2-1].Timestamp
			sortedTimestampTouchpoint2 := sortedLogs[numberOfLines/3-1].Timestamp
			sortedTimestampTouchpoint3 := sortedLogs[numberOfLines/1-1].Timestamp
			assert.True(t,
				(timestampTouchpoint0.Unix() != sortedTimestampTouchpoint0.Unix()) ||
					(timestampTouchpoint1.Unix() != sortedTimestampTouchpoint1.Unix()) ||
					(timestampTouchpoint2.Unix() != sortedTimestampTouchpoint2.Unix()) ||
					(timestampTouchpoint3.Unix() != sortedTimestampTouchpoint3.Unix()))

			assert.Exactly(t, &sortedLogs, &logs, "reference")
			assert.True(t, (&sortedTimestampTouchpoint0 != &timestampTouchpoint0) ||
				(&sortedTimestampTouchpoint1 != &timestampTouchpoint1) ||
				(&sortedTimestampTouchpoint2 != &timestampTouchpoint2) ||
				(&sortedTimestampTouchpoint3 != &timestampTouchpoint3))

			resultPath := filepath.Join(configuration.DirBinStorage, configuration.PathLogs, "test_result.log")
			err := LogWriteLines(resultPath, sortedLogs)
			assert.NoError(t, err)

			fd, err := os.Open(resultPath)
			assert.NoError(t, err)

			ops := 0
			reader := bufio.NewReader(fd)

			for {
				line, _, err := reader.ReadLine()
				if err != nil {
					break
				}
				ops += 1
				assert.NoError(t, err)

				if ops == 1 {
					assert.Exactly(t, detectTimestamp(line), sortedTimestampTouchpoint0)
				}
			}
			assert.Exactly(t, ops, numberOfLines)
			fd.Close()

			break
		}
	})
}
