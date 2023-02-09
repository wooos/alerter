package main

import (
	"github.com/spf13/cobra"
	"github.com/wooos/alerter/internal/config"
	"github.com/wooos/alerter/internal/server"
)

type options struct {
	configFile string
}

func (o *options) runCommand(cmd *cobra.Command, args []string) {
	config.LoadConfig(o.configFile)

	server.RunServer()
}

func NewCommand() *cobra.Command {
	opt := options{}
	cmd := &cobra.Command{
		Use: "alerter",
		Run: opt.runCommand,
	}

	flags := cmd.Flags()

	flags.StringVarP(&opt.configFile, "config", "c", "config.yaml", "config file path")

	return cmd
}

func main() {
	cmd := NewCommand()
	cmd.Execute()
}
