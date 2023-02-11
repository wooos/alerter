package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/wooos/alerter/internal/config"
	"github.com/wooos/alerter/internal/pkg/version"
	"github.com/wooos/alerter/internal/server"
)

type options struct {
	configFile  string
	showVersion bool
}

func (o *options) runCommand(cmd *cobra.Command, args []string) {
	if o.showVersion {
		fmt.Printf("%#v\n", version.Get())
		return
	}

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
	flags.BoolVarP(&opt.showVersion, "version", "v", false, "show version")

	return cmd
}

func main() {
	cmd := NewCommand()
	cmd.Execute()
}
