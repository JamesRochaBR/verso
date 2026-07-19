package cli

type Command interface {
	Name() string
	Run(args []string) error
}

var commands = []Command{
	VersionCommand{},
	InitCommand{},
	ValidateCommand{},
	InspectCommand{},
	BuildCommand{},
}
