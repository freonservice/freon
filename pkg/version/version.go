package version

import (
	"fmt"
	"runtime"
)

type Info struct {
	GitVersion string `json:"git_version"`
	BuildDate  string `json:"build_date"`
	GoVersion  string `json:"go_version"`
	Compiler   string `json:"compiler"`
	Platform   string `json:"platform"`
}

var (
	gitVersion = "NoExecutable"
	buildDate  = "NoExecutable"
	goVersion  = runtime.Version()
	compiler   = runtime.Compiler
	platform   = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)

func Get() Info {
	return Info{
		GitVersion: gitVersion,
		BuildDate:  buildDate,
		GoVersion:  goVersion,
		Compiler:   compiler,
		Platform:   platform,
	}
}
