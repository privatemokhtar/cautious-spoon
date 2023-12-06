package pet

import "github.com/df-mc/dragonfly/server/cmd"

type CreateCommand struct {
	// This sub command will be accessible as /pet create <name>.
	Create cmd.SubCommand `cmd:"create"`
	Name cmd.Varargs `cmd:"name"`
}

func (CreateCommand) Run(source cmd.Source, output *cmd.Output) {}

type DeleteCommand struct {
	// This subcommand will be accessible as /pet Delete.
	Delete cmd.SubCommand
}

func (DeleteCommand) Run(source cmd.Source, output *cmd.Output) {}
