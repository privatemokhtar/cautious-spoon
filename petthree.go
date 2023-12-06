package pet

import "github.com/df-mc/dragonfly/server/cmd"

func registerCommands() {
    // We pass both cmd.Runnable implementations here. Passing nil for the aliases is valid.
    cmd.Register(cmd.New("pet", "Pet management", nil, CreateCommand{}, DeleteCommand{}))
}
