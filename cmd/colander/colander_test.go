package main

import (
	"bufio"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/supermercato24/colander/config"
	"github.com/supermercato24/configuration"
)

const (
	numberOfLines0 = 5
	numberOfLines1 = 18
	dailyPattern0  = "d2-2017-12-11.log"
	dailyPattern1  = "d2-2017-12-11*.log"
	dailyLogFile0  = "d2-2017-12-11.log"
	dailyLogFile1  = "d2-2017-12-11-ws01.log"
	dailyLogFile2  = "d2-2017-12-11-ws02.log"
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

func withScreen(t *testing.T) {
	Colander(&ColanderOptions{
		dir:  filepath.Join(configuration.DirBinStorage, configuration.PathLogs),
		show: true,
	})
}

func withoutScreen(t *testing.T) {
	Colander(&ColanderOptions{
		dir:     filepath.Join(configuration.DirBinStorage, configuration.PathLogs),
		pattern: dailyPattern1,
	})

	resultPath := filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile0)

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
			assert.Exactly(t, []byte("[2017-12-01 23:16:20] d2-201712-11"), line)
		} else if ops == numberOfLines1 {
			assert.Exactly(t, []byte("[2017-12-01 23:26:22] d2-201712-10-ws01"), line)
		}
	}
	assert.Exactly(t, ops, numberOfLines1)
	fd.Close()
}

func withRemove(t *testing.T) {
	Colander(&ColanderOptions{
		dir:     filepath.Join(configuration.DirBinStorage, configuration.PathLogs),
		pattern: dailyPattern1,
		remove:  true,
	})

	resultPath := filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile0)

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
			assert.Exactly(t, []byte("[2017-12-01 23:16:20] d2-201712-11"), line)
		} else if ops == numberOfLines1 {
			assert.Exactly(t, []byte("[2017-12-01 23:26:22] d2-201712-10-ws01"), line)
		}
	}
	assert.Exactly(t, ops, numberOfLines1)
	fd.Close()

	_, err = os.Open(filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile1))
	assert.Error(t, err)
	_, err = os.Open(filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile2))
	assert.Error(t, err)
}

func singleWithoutScreen(t *testing.T) {
	Colander(&ColanderOptions{
		dir:     filepath.Join(configuration.DirBinStorage, configuration.PathLogs),
		pattern: dailyPattern0,
	})

	resultPath := filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile0)

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
			assert.Exactly(t, []byte("[2017-12-01 23:16:20] d2-201712-11"), line)
		} else if ops == numberOfLines0 {
			assert.Exactly(t, []byte("[2017-12-01 23:24:20] d2-201712-11"), line)
		}
	}
	assert.NotEqual(t, ops, numberOfLines1)
	assert.Exactly(t, ops, numberOfLines0)
	fd.Close()
}

func singleWithRemove(t *testing.T) {
	Colander(&ColanderOptions{
		dir:     filepath.Join(configuration.DirBinStorage, configuration.PathLogs),
		pattern: dailyPattern0,
		remove:  true,
	})

	resultPath := filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile0)

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
			assert.Exactly(t, []byte("[2017-12-01 23:16:20] d2-201712-11"), line)
		} else if ops == numberOfLines0 {
			assert.Exactly(t, []byte("[2017-12-01 23:24:20] d2-201712-11"), line)
		}
	}
	assert.NotEqual(t, ops, numberOfLines1)
	assert.Exactly(t, ops, numberOfLines0)
	fd.Close()

	fd, err = os.Open(filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile1))
	assert.NoError(t, err)
	fd.Close()
	fd, err = os.Open(filepath.Join(configuration.DirBinStorage, configuration.PathLogs, dailyLogFile2))
	assert.NoError(t, err)
	fd.Close()
}

func TestLog(t *testing.T) {
	assert.Exactly(t, config.Name, config.Name, "load init")

	//assert.True(t, t.Run("withScreen", withScreen))

	assert.True(t, t.Run("setUp0", func(t *testing.T) { logSetUp(t, 0) }))
	assert.True(t, t.Run("setUp1", func(t *testing.T) { logSetUp(t, 1) }))
	assert.True(t, t.Run("setUp2", func(t *testing.T) { logSetUp(t, 2) }))

	assert.True(t, t.Run("withoutScreen", withoutScreen))
	assert.True(t, t.Run("setUp0Refresh", func(t *testing.T) { logSetUp(t, 0) }))

	assert.True(t, t.Run("withRemove", withRemove))
	assert.True(t, t.Run("setUp0Restore", func(t *testing.T) { logSetUp(t, 0) }))
	assert.True(t, t.Run("setUp1Restore", func(t *testing.T) { logSetUp(t, 1) }))
	assert.True(t, t.Run("setUp2Restore", func(t *testing.T) { logSetUp(t, 2) }))

	assert.True(t, t.Run("singleWithoutScreen", singleWithoutScreen))

	assert.True(t, t.Run("singleWithRemove", singleWithRemove))
	assert.True(t, t.Run("setUp0RestoreSingle", func(t *testing.T) { logSetUp(t, 0) }))
}
