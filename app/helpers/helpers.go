package helpers

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	// basePathCache stores the value of BasePath after its first call.
	basePathCache string
)

// BasePath fetches project path, caching the result.
func BasePath() string {
	if basePathCache == "" {
		basePathCache = evaluateBasePath()
	}
	return basePathCache
}

// evaluateBasePath evaluates the base path for project.
func evaluateBasePath() string {
	if projectPath := os.Getenv("PROJECT_PATH"); projectPath != "" {
		return projectPath
	}
	if gopath := os.Getenv("GOPATH"); gopath != "" {
		gopath = filepath.Join(strings.Split(gopath, string(os.PathListSeparator))[0], "src")
		return gopath + "/github.com/Oleg-Smal-git/boosters-trial"
	}
	return "."
}
