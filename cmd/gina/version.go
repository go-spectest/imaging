package main

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var (
	// Version value is set by ldflags
	Version string //nolint
	// Name is cli command name
	Name = "imaging" //nolint
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "version",
		Short:   "Show " + Name + " command version information",
		Example: "   gina version",
		Run:     version,
	}
}

// version return reddit-downloader command version.
// Version global variable is set by ldflags.
func version(_ *cobra.Command, _ []string) {
	version := "unknown"
	if Version != "" {
		version = Version
	} else if buildInfo, ok := debug.ReadBuildInfo(); ok {
		version = buildInfo.Main.Version
	}
	fmt.Printf("%s version %s (under MIT LICENSE)\n", Name, version)
}
