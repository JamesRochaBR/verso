package cli

import "fmt"

type InitCommand struct{}

func (InitCommand) Name() string {
	return "init"
}

func (InitCommand) Run(args []string) error {
	fmt.Println("init not implemented")
	return nil
}