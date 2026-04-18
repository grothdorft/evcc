package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	// Version is the application version, set during build
	Version = "0.0.1-dev"
	// Commit is the git commit hash, set during build
	Commit = "unknown"
	// Date is the build date, set during build
	Date = "unknown"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information",
	Long:  `Print the version, commit hash, and build date of this evcc binary.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("evcc version %s\n", Version)
		fmt.Printf("  commit: %s\n", Commit)
		fmt.Printf("  built:  %s\n", Date)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
