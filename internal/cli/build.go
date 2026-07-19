package cli

import "fmt"

type BuildCommand struct{}

func (BuildCommand) Name() string {
	return "build"
}

func (BuildCommand) Run(args []string) error {
	fmt.Println("build not implemented")
	return nil
}