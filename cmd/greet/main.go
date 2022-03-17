package main

import (
	"github.com/urfave/cli/v2"
	"goapplayout/greet"
	"log"
	"os"
)

func main() {
	var language string

	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Aliases:     []string{"l"},
				Value:       "english",
				Usage:       "language for the greeting",
				Required:    true,
				Destination: &language,
			},
		},
		Name:  "pkg",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			name := "Nefertiti"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			greet.SayHello(language, name)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
