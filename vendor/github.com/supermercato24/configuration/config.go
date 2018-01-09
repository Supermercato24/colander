// Package configuration implements methods for get config variables.
//
// Config main project.
// jenkins/gdo/compiled = go main.
// jenkins/gdo/.env.
// jenkins/gdo/config/soa.php.
package configuration

import (
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

const (
	PathBin     = "compiled" // go code
	PathScript  = "script"   // python code
	PathFile    = "file"     // file          for python and php code
	PathLib     = "lib"      // lib           for python and php code
	PathApp     = "app"      // app           for python and php code
	PathConfig  = "config"   // configuration for python and php code
	PathStorage = "storage"  // storage       for python and php code
	PathLogs    = "logs"     // log           for python and php code
)

var (
	DirCwd            string // current working dir
	DirProject        string // php framework
	DirProjectApp     string // php app dir
	DirProjectConfig  string // php config dir
	DirProjectStorage string // php storage dir
	DirScript         string // python framework
	DirScriptApp      string // python app dir
	DirScriptConfig   string // python config dir
	DirScriptStorage  string // python storage dir
	DirBinStorage     string // go storage dir
)

func init() {
	// get current working directory
	cwd, _ := os.Getwd()
	cwd, _ = filepath.EvalSymlinks(cwd)

	relativeDir := ".."

	switch baseDir := filepath.Base(cwd); baseDir {
	case PathBin:
		// pass
	case PathConfig, PathFile, PathLib:
		relativeDir = filepath.Join(relativeDir, "..")
	default:
		relativeDir = filepath.Join(relativeDir, "..", "..")
	}

	projectDir, _ := filepath.Abs(filepath.Join(cwd, relativeDir))

	DirCwd = cwd            // /
	DirProject = projectDir // gdo/
	BuildProject()
}

// BuildProject (or rebuild) dir structure
func BuildProject() {
	DirProjectApp = filepath.Join(DirProject, PathApp)              // gdo/app
	DirProjectConfig = filepath.Join(DirProject, PathConfig)        // gdo/config
	DirProjectStorage = filepath.Join(DirProject, PathStorage)      // gdo/storage
	DirScript = filepath.Join(DirProject, PathScript)               // gdo/script
	DirScriptApp = filepath.Join(DirScript, PathApp)                // gdo/script/app
	DirScriptConfig = filepath.Join(DirScriptApp, PathConfig)       // gdo/script/app/config
	DirScriptStorage = filepath.Join(DirScript, PathStorage)        // gdo/script/storage
	DirBinStorage = filepath.Join(DirProject, PathBin, PathStorage) // compiled/script/storage
}

// Read a file, from a path and extract configuration from regexp.
func cfg(filePath string, regexpString string) (object map[string][]byte, err error) {
	// content of file
	var bytesContent []byte

	isAbs := path.IsAbs(filePath)

	// absolute path
	if isAbs {
		bytesContent, err = ioutil.ReadFile(filePath)
	}

	// directory path
	if !isAbs && err == nil {
		bytesContent, err = ioutil.ReadFile(filepath.Join(DirProject, filePath))
	}

	// relative path
	if err != nil {
		bytesContent, err = ioutil.ReadFile(filePath)
	}

	// open and read an absolute file
	if err != nil {
		return
	}

	// compiled regex
	re := regexp.MustCompile(regexpString)
	names := re.SubexpNames()
	subexpressions := re.NumSubexp()
	matches := re.FindAllSubmatch(bytesContent, -1)
	if matchesLen := len(matches); matchesLen < subexpressions {
		subexpressions = matchesLen
	}

	// temporary dict of results
	object = map[string][]byte{}
	for i, key := 0, 1; i < subexpressions; i, key = i+1, key+1 {
		match := matches[i]
		name := names[key]
		//for key, match := range matches[i] {

		//if len(match) > 0 && len(name) > 0 {
		object[name] = match[key]
		//}

		//}
	}

	return
}
