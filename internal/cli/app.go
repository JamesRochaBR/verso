package cli

import "fmt"

func Run(args []string) int {
	if len(args) < 2 {
		fmt.Println("Verso CLI")
		return 0
	}

	for _, command := range commands {
		if command.Name() == args[1] {
			if err := command.Run(args[2:]); err != nil {
				fmt.Println("Error:", err)
				return 1
			}
			return 0
		}
	}

	fmt.Println("Unknown command:", args[1])
	return 1
}
