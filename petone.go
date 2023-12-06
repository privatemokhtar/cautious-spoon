package pet

import "github.com/df-mc/dragonfly/server/cmd"

type CreateCommand struct {
	Name cmd.Varargs `cmd:"name"`
}

func (CreateCommand) Run(source cmd.Source, output *cmd.Output) {}

type DeleteCommand struct {}

func (DeleteCommand) Run(source cmd.Source, output *cmd.Output) {}
