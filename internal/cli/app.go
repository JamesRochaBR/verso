package cli

import "fmt"

func Run(args []string) int {
	switch {
	case len(args) >= 2 && args[1] == "version":
		return runVersion()
	default:
		fmt.Println("Verso CLI")
		return 0
	}
}