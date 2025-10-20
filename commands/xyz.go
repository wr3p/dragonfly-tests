package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type XYZ struct{}

var opened = map[string]bool{}

func (cmd XYZ) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	p, ok := source.(*player.Player)
	if !ok {
		output.Print("You can only use this command in-game!")
		return
	}

	id := p.Name()
	opened[id] = !opened[id]

	if opened[id] {
		p.ShowCoordinates()
		output.Print("§aCoordinates are displayed.")
	} else {
		p.HideCoordinates()
		output.Print("§cCoordinates are hidden.")
	}
}
