package cli

import "fmt"

func Run(args []string) int {
	if len(args) < 2 {
		printHelp()
		return 0
	}

	switch args[1] {
	case "--help", "-h", "help":
		printHelp()
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
	fmt.Println()
	printHelp()

	return 1
}
