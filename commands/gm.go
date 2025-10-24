package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"

	"strings"
)

type GAMEMODE struct {
	Mode cmd.Optional[string] `cmd:"mode"`
}

func (cmd GAMEMODE) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	p, ok := source.(*player.Player)
	if !ok {
		output.Print("You can only use this command in-game!")
		return
	}

	modeArg, ok := cmd.Mode.Load()
	if !ok {
		p.Message("§cUsage: /gamemode <mode>")
		return
	}

	modeArg = strings.ToLower(modeArg)

	var mode world.GameMode
	var mode_name string
	switch modeArg {
	case "0", "s", "survival":
		mode = world.GameModeSurvival
		mode_name = "Survival"
	case "1", "c", "creative":
		mode = world.GameModeCreative
		mode_name = "Creative"
	case "2", "a", "adventure":
		mode = world.GameModeAdventure
		mode_name = "Adventure"
	case "3", "sp", "spectator":
		mode = world.GameModeSpectator
		mode_name = "Spectator"
	default:
		p.Message("§cInvalid gamemode! Please use: 0/1/2/3 or s/c/a/sp")
		return
	}

	p.SetGameMode(mode)
	p.Messagef("§aYour gamemode has been set to §e%s", mode_name)
}
