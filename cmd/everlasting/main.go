package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

// github

func main() {
	app := &cli.App{
		Name:        "everlasting",
		Usage:       "everlasting [sub commands] [flags]",
		Description: "This is simple setup tool.",
		Commands: []*cli.Command{
			tuiCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
