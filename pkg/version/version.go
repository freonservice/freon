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
)

func Get() Info {
	return Info{
		GitVersion: gitVersion,
		BuildDate:  buildDate,
		GoVersion:  runtime.Version(),
		Compiler:   runtime.Compiler,
		Platform:   fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}
