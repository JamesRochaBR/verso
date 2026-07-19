package cli

import "fmt"

const Version = "0.2.0-dev"

func runVersion() int {
	fmt.Printf("Verso %s\n", Version)
	return 0
}