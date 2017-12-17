// Package main implements methods to colander log aggregator.
//
// Command.
package main

import (
	"fmt"
	"os"
	"path/filepath"

	"supermercato24.it/colander/file"
)

type ColanderOptions struct {
	dir     string
	pattern string
	remove  bool
	show    bool
}

func Colander(options *ColanderOptions) {
	var showClosure func(resultPath string, logs file.Aggregate)
	if options.show {
		showClosure = func(_ string, logs file.Aggregate) {
			for _, log := range logs {
				fmt.Println(string(log.Body))
			}
		}
	} else {
		showClosure = func(resultPath string, logs file.Aggregate) {
			err := file.LogWriteLines(resultPath, logs)
			if err != nil {
				panic(err)
			}
		}
	}

	var removeClosure func(resultPath string, paths []string)
	if options.remove {
		removeClosure = func(resultPath string, paths []string) {
			_, pathResult := filepath.Split(resultPath)
			for _, path := range paths {
				_, pathFile := filepath.Split(path)
				if pathResult != pathFile { // don't remove result file
					err := os.Remove(path)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	} else {
		removeClosure = func(_ string, _ []string) {}
	}

	file.Glob(options.dir, options.pattern, func(matches *file.GlobMatches) {
		for _, match := range matches.Files {
			category := match.Category
			for _, match := range match.Logs {
				date := match.Time.Format("2006-01")

				logs := file.LogReadLines(match.Paths)
				file.SortByTimestamp(logs)

				var resultPath string
				if !options.show {
					resultPath = fmt.Sprintf("%s-%s", category, date)
					if match.Day != int64(0) {
						resultPath = fmt.Sprintf("%s-%d", resultPath, match.Time.Day())
					}
					resultPath = fmt.Sprintf("%s%s", resultPath, file.LogExtension)
					resultPath = filepath.Join(options.dir, resultPath)
				}

				showClosure(resultPath, logs)
				if len(match.Paths) > 1 {
					removeClosure(resultPath, match.Paths)
				}
			}
		}
	})
}
