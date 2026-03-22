package cmd

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"runtime/debug"
	"time"

	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var (
	Version string
	Commit  string
	Date    string
)

func initBuildInfo() {
	buildInfo, ok := debug.ReadBuildInfo()
	if !ok {
		return
	}

	if Version == "" {
		if buildInfo.Main.Version != "" && buildInfo.Main.Version != "(devel)" {
			Version = buildInfo.Main.Version
		}
	}

	if Commit == "" {
		Commit = buildSetting(buildInfo, "vcs.revision")
	}

	if Date == "" {
		Date = buildSetting(buildInfo, "vcs.time")
	}
}

func buildSetting(buildInfo *debug.BuildInfo, key string) string {
	for _, setting := range buildInfo.Settings {
		if setting.Key == key {
			return setting.Value
		}
	}

	return ""
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "simplerpc version",
	Long:  `simplerpc version`,
	Run: func(cmd *cobra.Command, args []string) {
		printVersion()
	},
}

func printVersion() {
	var versionBuffer bytes.Buffer
	resolvedVersion := Version
	resolvedCommit := Commit
	resolvedDate := Date

	if resolvedVersion == "" {
		resolvedVersion = "unknown"
	}
	versionBuffer.WriteString(fmt.Sprintf("simplerpc version %s %s/%s\n", resolvedVersion, runtime.GOOS, runtime.GOARCH))

	versionBuffer.WriteString(fmt.Sprintf("Go version %s\n", runtime.Version()))

	if resolvedCommit == "" {
		resolvedCommit = "unknown"
	}
	versionBuffer.WriteString(fmt.Sprintf("Git commit %s\n", resolvedCommit))

	if resolvedDate != "" {
		resolvedDate = cast.ToString(cast.ToTimeInDefaultLocation(resolvedDate, time.Local))
	} else {
		resolvedDate = "unknown"
	}
	versionBuffer.WriteString(fmt.Sprintf("Build date: %s\n", resolvedDate))

	fmt.Print(versionBuffer.String())
}

func init() {
	initBuildInfo()
	_ = os.Setenv("VERSION", Version)
	_ = os.Setenv("COMMIT", Commit)
	_ = os.Setenv("DATE", Date)

	rootCmd.AddCommand(versionCmd)
}
