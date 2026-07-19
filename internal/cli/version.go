package cli

import "fmt"

const Version = "0.2.0-dev"

type VersionCommand struct{}

func (VersionCommand) Name() string {
	return "version"
}

func (VersionCommand) Run(args []string) error {
	fmt.Printf("Verso %s\n", Version)
	return nil
}
