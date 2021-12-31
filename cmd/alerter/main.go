package main

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/wooos/alerter/internal/config"
	"github.com/wooos/alerter/internal/server"
)

type options struct {
	configFile string
}

func (o *options) runCommand(cmd *cobra.Command, args []string) {
	conf, err := config.LoadConfig(o.configFile)
	if err != nil {
		log.Fatalln(err)
	}

	server.RunServer(conf)
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
