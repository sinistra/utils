package utils

import (
	"fmt"
	"runtime/debug"
)

// PrintVersion prints the application version
func PrintVersion() string {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return fmt.Sprintln("Unable to determine version information.")
	}

	version := "Version: unknown"
	if buildInfo.Main.Version != "" {
		version = fmt.Sprintf("Version: %s\n", buildInfo.Main.Version)
	}
	return version
}
