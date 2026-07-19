package cli

import (
	"fmt"

	"github.com/james-rocha/verso/internal/build"
)

type BuildCommand struct{}

func (BuildCommand) Name() string {
	return "build"
}

func (BuildCommand) Run(args []string) error {

	if len(args) == 0 {
		return fmt.Errorf("missing project path")
	}

	return build.Build(args[0])
}
