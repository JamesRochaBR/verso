package cli

import (
	"fmt"
)

func Run(args []string) int {
	fmt.Println("Verso CLI")

	commands := []Command{
		InitCommand{},
		ValidateCommand{},
		BuildCommand{},
	}

	_ = commands

	return 0
}