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
				Name:     "config",
				Aliases:  []string{"c"},
				Usage:    "Load configuration from `FILE`",
				Required: true,
			},
			&cli.StringFlag{
				Name:    "lang",
				Aliases: []string{"l"},
				Value:   "english",
				Usage:   "language for the greeting",
				//Required:    true,
				Destination: &language,
			},
		},
		Name:  "greet",
		Usage: "fight the loneliness!",
		Action: func(c *cli.Context) error {
			name := "User"
			if c.NArg() > 0 {
				name = c.Args().Get(0)
			}
			configPath := c.String("config")
			config, err := greet.NewConfig(configPath)
			if err != nil {
				log.Fatal(err)
			}
			greet.SayHello(config, language, name)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
