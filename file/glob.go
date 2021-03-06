// Package file implements methods for get logs from files.
//
// Log Finding.
package file

import (
	"fmt"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	splitAfter  = 2
	globPattern = "*-[0-9][0-9][0-9][0-9]*-*"
)

// GlobMatches map of Files
type GlobMatches struct {
	Files map[string]*FileMatches
}

// FileMatches map of Logs with same Category.
type FileMatches struct {
	Category string
	Logs     map[time.Time]*LogMatches
}

// LogMatches single line of log with timestamp.
type LogMatches struct {
	Year  int64
	Month int64
	Day   int64
	Time  time.Time
	Paths []string
}

func splitLogFilesByName(pattern string) *GlobMatches {
	var matches []string
	globMatches := &GlobMatches{
		Files: make(map[string]*FileMatches),
	}

	matches, _ = filepath.Glob(pattern)
	for _, match := range matches {
		matchWithoutExtension := strings.TrimSuffix(match, filepath.Ext(LogExtension))

		_, file := filepath.Split(matchWithoutExtension)
		title := strings.SplitN(file, "-", splitAfter)
		if len(title) < splitAfter {
			continue
		}

		dates := strings.SplitN(title[1], "-", splitAfter*2)
		if len(dates) < splitAfter {
			continue
		}

		fileMatches := &FileMatches{
			Category: title[0],
		}
		if _, ok := globMatches.Files[fileMatches.Category]; !ok {
			fileMatches.Logs = make(map[time.Time]*LogMatches)
			globMatches.Files[fileMatches.Category] = fileMatches
		}
		fileMatches = globMatches.Files[fileMatches.Category]

		logMatches := &LogMatches{}
		logMatches.Year, _ = strconv.ParseInt(dates[0], 10, 32)
		logMatches.Month, _ = strconv.ParseInt(dates[1], 10, 32)
		if len(dates) >= splitAfter+1 {
			logMatches.Day, _ = strconv.ParseInt(dates[2], 10, 32)
		}
		logMatches.Time = time.Date(
			int(logMatches.Year), time.Month(logMatches.Month), int(logMatches.Day),
			0, 0, 0, 0, time.UTC)
		if _, ok := fileMatches.Logs[logMatches.Time]; !ok {
			logMatches.Paths = make([]string, 0)
			fileMatches.Logs[logMatches.Time] = logMatches
		}
		logMatches = fileMatches.Logs[logMatches.Time]

		logMatches.Paths = append(logMatches.Paths, match)
	}

	return globMatches
}

// Glob search files into directory by pattern.
func Glob(dirPath string, pattern string, f func(matches *GlobMatches)) {
	// pay_debug-2017-11-15-ws02.log
	// pay_debug-2017-11-15-ws01.log
	// pay_debug-2017-11-15-ws03.log
	// pay_debug-2017-11-16-ws02.log
	// pay_debug-2017-11-16-ws01.log
	// pay_debug-2017-11-16-ws03.log
	// sms-2017-11-ws01.log
	// sms-2017-11-ws02.log
	// sms-2017-11-ws03.log
	// sms-2017-12-ws01.log
	// sms-2017-12-ws02.log
	// sms-2017-12-ws03.log

	// {"pay_debug": {"2017-11-15":["pay_debug-2017-11-15-ws01.log","pay_debug-2017-11-15-ws02.log"]}}

	if pattern == "" {
		pattern = fmt.Sprintf("%s%s", globPattern, LogExtension)
	}
	if dirPath != "" {
		pattern = fmt.Sprintf("%s/%s", filepath.Clean(dirPath), pattern)
	}

	f(splitLogFilesByName(pattern))
}
