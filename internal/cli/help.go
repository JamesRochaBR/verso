package cli

import "fmt"

func printHelp() {
	fmt.Println(`Verso

Usage:
  verso <command> [arguments]

Commands:
  init        Create a new Verso project
  validate    Validate a project
  inspect     Inspect project components
  build       Build project context
  prompt      Generate AI prompt
  version     Show version

Use "verso <command> --help" for more information about a command.`)
}
