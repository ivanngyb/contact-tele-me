package main

import (
	"log"
	"os"
	"server/api"

	"github.com/urfave/cli/v2"
)

const envPrefix = "IVANAPI"

func main() {
	app := &cli.App{
		Name:  "Ivan's Portfolio API Server",
		Usage: "API Server to showcase features on portfolio",
		Commands: []*cli.Command{
			{
				Name:    "server",
				Aliases: []string{"s"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "api_addr", Value: ":8084", EnvVars: []string{envPrefix + "_API_ADDR", "API_ADDR"}, Usage: "Port for which server will run on"},
				},
				Usage: "Run server",
				Action: func(ctx *cli.Context) error {
					apiConfig := &api.Config{
						Address: ctx.String("api_addr"),
					}

					api := api.NewAPI(ctx.Context, apiConfig)

					err := api.Run(ctx.Context)
					if err != nil {
						log.Fatal("fail to start api")
					}

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
