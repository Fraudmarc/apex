// Package version oututs the version.
package version

import (
	"fmt"

	"github.com/tj/cobra"

	"github.com/fraudmarc/apex/cmd/apex/root"
)

// Version of program.
const Version = "v1.0.0-rc5"

// Command config.
var Command = &cobra.Command{
	Use:              "version",
	Short:            "Print version of Apex",
	PersistentPreRun: root.PreRunNoop,
	Run:              run,
}

// Initialize.
func init() {
	root.Register(Command)
}

// Run command.
func run(c *cobra.Command, args []string) {
	fmt.Printf("Apex version %s\n", Version)
}
