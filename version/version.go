package version

import "fmt"

var AppVersion string


var GitCommit string

func VersionInfo() string {
	return fmt.Sprintf("version: %s\ncommit: %s", AppVersion, GitCommit)
}
