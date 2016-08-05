package main

import (
	"fmt"
	"os"

	"github.com/bsdpunk/beastietools/command"
	"github.com/codegangsta/cli"
)

var GlobalFlags = []cli.Flag{}

var Commands = []cli.Command{
	{
		Name:   "extip",
		Usage:  "",
		Action: command.CmdExtip,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "arp",
		Usage:  "",
		Action: command.CmdArp,
		Flags:  []cli.Flag{},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
