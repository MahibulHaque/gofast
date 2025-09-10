package cmd

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/spf13/cobra"
)

var GofastVersion string

// Go Fast needs to be built in a specific way to provide useful version information.
// First we try to get the version from ldflags embedded into GoFastVersion.
// Then we try to get the version from from the go.mod build info.
// If Go Fast is installed with a specific version tag or using @latest then that version will be included in bi.Main.Version.
// This won't give any version info when running 'go install' with the source code locally.
// Finally we try to get the version from other embedded VCS info.
func getGoFastVersion() string {
	noVersionAvailable := "No version info available for this build, run 'gofast version' for additional info"

	if len(GofastVersion) != 0 {
		return GofastVersion
	}

	bi, ok := debug.ReadBuildInfo()
	if !ok {
		return noVersionAvailable
	}

	// If no main version is available, Go defaults it to (devel)
	if bi.Main.Version != "(devel)" {
		return bi.Main.Version
	}

	var vcsRevision string
	var vcsTime time.Time
	for _, setting := range bi.Settings {
		switch setting.Key {
		case "vcs.revision":
			vcsRevision = setting.Value
		case "vcs.time":
			vcsTime, _ = time.Parse(time.RFC3339, setting.Value)
		}
	}

	if vcsRevision != "" {
		return fmt.Sprintf("%s, (%s)", vcsRevision, vcsTime)
	}

	return noVersionAvailable
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Display application version information.",
	Long: `
The version command provides information about the application's version.

Go Fast requires version information to be embedded at compile time.
For detailed version information, Go Fast needs to be built as specified in the README installation instructions.
If Go Fast is built within a version control repository and other version info isn't available,
the revision hash will be used instead.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		version := getGoFastVersion()
		fmt.Printf("Go Fast CLI version: %v\n", version)
	},
}
