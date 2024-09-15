package main

import "fmt"

var (
	version string
	commit  string
	date    string
	builtBy string
)

func getVersion() string {
	return fmt.Sprintf("\nVersion\t: %s\nCommit\t: %s\nDate\t: %s\nBuilt By: %s\n", version, commit, date, builtBy)
}
