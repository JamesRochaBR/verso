package main

import (
	"os"

	"github.com/james-rocha/verso/internal/cli"
)

func main() {
	os.Exit(cli.Run(os.Args))
}