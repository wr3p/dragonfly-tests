package main

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"

	"github.com/wr3p/dragonfly-tests/commands"
)

func main() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
	chat.Global.Subscribe(chat.StdoutSubscriber{})
	conf, err := readConfig(slog.Default())
	if err != nil {
		panic(err)
	}

	srv := conf.New()
	// commands etc
	registerCmds()

	// others
	srv.CloseOnProgramEnd()

	srv.Listen()
	for joined := range srv.Accept() {
		worldTx := joined.Tx()
		for other := range srv.Players(worldTx) {
			if other.UUID() != joined.UUID() {
				other.Message(fmt.Sprintf("%s joined the world!", joined.Name())) // not sure of this
			}
			joined.Message("Welcome!")
		}
	}

	/*for p := range srv.Accept() {
		_ = p
	}*/
}

func registerCmds() {
	cmd.Register(cmd.New("xyz", "Shows/hides coordinates", []string{"coordinates"}, commands.XYZ{}))
	cmd.Register(cmd.New("transfer", "Transfer yourself to another server", []string{"go"}, commands.TRANSFER{}))
}

// readConfig reads the configuration from the config.toml file, or creates the
// file if it does not yet exist.
func readConfig(log *slog.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
