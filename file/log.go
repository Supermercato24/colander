package file

import (
	"bufio"
	"os"
	"regexp"
	"time"
)

const (
	LogExtension = ".log"
)

type Log struct {
	Timestamp time.Time
	Body      []byte
}

var (
	timestampRegex = regexp.MustCompile(`[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}`)
	newLine        = []byte("\n")
)

// Try to detect timestamp from body with regex and time parser
// [2017-12-01 23:50:07]
// [2017-12-01 23:50:10]
func detectTimestamp(line []byte) (timestampDetected time.Time) {

	timestamp := timestampRegex.Find(line)
	if len(timestamp) > 0 {
		timestampDetected, _ = time.Parse("2006-01-02 15:04:05", string(timestamp))
	}

	return
}

func LogReadLines(logFiles []string) []Log {
	var logs []Log

	for _, logFile := range logFiles {
		fd, err := os.Open(logFile)
		if err != nil {
			panic(err)
		}
		defer fd.Close()

		reader := bufio.NewReader(fd)

		for {
			line, _, err := reader.ReadLine()
			if err != nil {
				break
			}

			timestamp := detectTimestamp(line)
			if (time.Time{}) == timestamp {
				continue
			}
			log := Log{
				Timestamp: timestamp,
				Body:      line,
			}
			logs = append(logs, log)
		}
	}

	return logs
}

func LogWriteLines(path string, logs []Log) (err error) {
	os.Remove(path)
	fd, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
	defer fd.Close()

	writer := bufio.NewWriter(fd)

	for _, log := range logs {
		_, err = writer.Write(log.Body)
		if err != nil {
			return
		}
		writer.Write(newLine)
	}
	writer.Flush()

	return
}
