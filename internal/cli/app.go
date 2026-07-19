package cli

import "fmt"

func Run(args []string) int {
	if len(args) < 2 {
		fmt.Println("Verso CLI")
		return 0
	}

	switch args[1] {

	case "version":
		return runVersion()

	case "init":
		cmd := InitCommand{}

		if err := cmd.Run(args[2:]); err != nil {
			fmt.Println("Error:", err)
			return 1
		}

		return 0

	case "validate":
		cmd := ValidateCommand{}

		if err := cmd.Run(args[2:]); err != nil {
			fmt.Println("Error:", err)
			return 1
		}

		return 0

	case "inspect":
		cmd := InspectCommand{}

		if err := cmd.Run(args[2:]); err != nil {
			fmt.Println("Error:", err)
			return 1
		}

		return 0

	default:
		fmt.Println("Unknown command:", args[1])
		return 1
	}
}