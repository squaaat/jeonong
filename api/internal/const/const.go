package _const

import (
	"runtime"
	"strings"
)

const (
	Project = "nearsfeed"
	App     = "nearsfeed-api"

	KeyEnv  = "J_ENV"
	KeyCicd = "J_CICD"

	EnvAlpha = "alpha"
	EnvProd  = "prod"
)

var (
	ProjectRootPath string
)

func init() {
	_, constFilePath, _, _ := runtime.Caller(0)
	splits := strings.Split(constFilePath, "/")
	absolutePathSplits := splits[:(len(splits) - 3)]
	ProjectRootPath = strings.Join(absolutePathSplits, "/")
}
