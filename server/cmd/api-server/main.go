package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "Ivan's Portfolio API Server",
		Usage: "API Server to showcase features on portfolio",
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"s"},
				Flags:   []cli.Flag{},
				Usage:   "Run server",
				Action: func(ctx *cli.Context) error {
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err.Error())
	}
}
