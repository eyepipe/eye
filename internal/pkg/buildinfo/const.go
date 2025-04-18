package buildinfo

import (
	"fmt"
	"runtime"
)

var (
	BuildArgVersion   string
	BuildArgGitCommit string
	BuildArgTime      string
)

// go tool nm <your binary> | grep <your variable>
func FullVersionString() string {
	return fmt.Sprintf(
		"%s on %s_%s, compiled by %s, commit %s, time %s",
		BuildArgVersion,
		runtime.GOOS,
		runtime.GOARCH,
		runtime.Version(),
		BuildArgGitCommit,
		BuildArgTime,
	)
}
