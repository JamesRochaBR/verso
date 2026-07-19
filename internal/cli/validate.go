package cli

import "fmt"

type ValidateCommand struct{}

func (ValidateCommand) Name() string {
	return "validate"
}

func (ValidateCommand) Run(args []string) error {
	fmt.Println("validate not implemented")
	return nil
}