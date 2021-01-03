package main

import (
	"github.com/hachi-n/everlasting"
	"github.com/urfave/cli/v2"
)

func tuiCommand() *cli.Command {
	return &cli.Command{
		Name:        "tui",
		Usage:       "everlasting tui",
		Description: "simple tui.",
		Flags: []cli.Flag{
		},
		Action: func(c *cli.Context) error {
			return everlasting.TuiStart("/resources", true)
		},
	}
}
