package main

import (
	"os"

	"github.com/apex/log"
	"github.com/apex/log/handlers/cli"

	"github.com/fraudmarc/apex/cmd/apex/root"

	// commands
	_ "github.com/fraudmarc/apex/cmd/apex/alias"
	_ "github.com/fraudmarc/apex/cmd/apex/autocomplete"
	_ "github.com/fraudmarc/apex/cmd/apex/build"
	_ "github.com/fraudmarc/apex/cmd/apex/delete"
	_ "github.com/fraudmarc/apex/cmd/apex/deploy"
	_ "github.com/fraudmarc/apex/cmd/apex/docs"
	_ "github.com/fraudmarc/apex/cmd/apex/exec"
	_ "github.com/fraudmarc/apex/cmd/apex/infra"
	_ "github.com/fraudmarc/apex/cmd/apex/init"
	_ "github.com/fraudmarc/apex/cmd/apex/invoke"
	_ "github.com/fraudmarc/apex/cmd/apex/list"
	_ "github.com/fraudmarc/apex/cmd/apex/logs"
	_ "github.com/fraudmarc/apex/cmd/apex/metrics"
	_ "github.com/fraudmarc/apex/cmd/apex/rollback"
	_ "github.com/fraudmarc/apex/cmd/apex/upgrade"
	_ "github.com/fraudmarc/apex/cmd/apex/version"

	// plugins
	_ "github.com/fraudmarc/apex/plugins/clojure"
	_ "github.com/fraudmarc/apex/plugins/golang"
	_ "github.com/fraudmarc/apex/plugins/hooks"
	_ "github.com/fraudmarc/apex/plugins/inference"
	_ "github.com/fraudmarc/apex/plugins/java"
	_ "github.com/fraudmarc/apex/plugins/nodejs"
	_ "github.com/fraudmarc/apex/plugins/python"
	_ "github.com/fraudmarc/apex/plugins/ruby"
	_ "github.com/fraudmarc/apex/plugins/rust_gnu"
	_ "github.com/fraudmarc/apex/plugins/rust_musl"
	_ "github.com/fraudmarc/apex/plugins/shim"
)

// Terraform commands.
var tf = []string{
	"apply",
	"destroy",
	"get",
	"graph",
	"init",
	"output",
	"plan",
	"refresh",
	"remote",
	"show",
	"taint",
	"untaint",
	"validate",
	"version",
}

// TODO(tj): remove this evil hack, necessary for now for cases such as:
//
//   $ apex --env prod infra deploy
//
// instead of:
//
//   $ apex infra --env prod deploy
//
func endCmdArgs(args []string, off int) []string {
	return append(args[0:off], append([]string{"--"}, args[off:]...)...)
}

func indexOf(args []string, key string) int {
	for i, arg := range args {
		if arg == key {
			return i
		}
	}
	return -1
}

func main() {
	log.SetHandler(cli.Default)

	args := os.Args[1:]

	// Cobra does not (currently) allow us to pass flags for a sub-command
	// as if they were arguments, so we inject -- here after the first TF command.
	// TODO(tj): replace with a real solution and send PR to Cobra #251
	if len(os.Args) > 1 && indexOf(os.Args, "infra") > -1 {
		off := 1

	out:
		for i, a := range args {
			for _, cmd := range tf {
				if a == cmd {
					off = i
					break out
				}
			}
		}

		args = endCmdArgs(args, off)
	} else if len(os.Args) > 1 && indexOf(os.Args, "exec") > -1 {
		args = endCmdArgs(args, indexOf(os.Args, "exec"))
	}

	root.Command.SetArgs(args)

	if err := root.Command.Execute(); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
