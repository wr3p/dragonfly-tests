package commands

import (
	"fmt"

	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type TRANSFER struct {
	Ip   cmd.Optional[string] `cmd:"ip"`
	Port cmd.Optional[string] `cmd:"port"`
}

func (cmd TRANSFER) Run(source cmd.Source, output *cmd.Output, tx *world.Tx) {
	p, ok := source.(*player.Player)
	if !ok {
		output.Print("You can only use this command in-game!")
		return
	}

	ip, ipOk := cmd.Ip.Load()
	port, portOk := cmd.Port.Load()

	if !ipOk || !portOk {
		p.Message("§cUsage: /transfer <ip> <port>")
		return
	}

	address := fmt.Sprintf("%s:%s", ip, port)
	if err := p.Transfer(address); err != nil {
		p.Messagef("§cFailed to transfer to %s: %v", address, err)
		return
	}

	p.Messagef("§aSuccessfully transferring to §e%s", address)
}
