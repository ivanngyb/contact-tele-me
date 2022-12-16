package main

import (
	"log"
	"os"
	"server/api"
	"server/logger"

	"github.com/urfave/cli/v2"
)

const envPrefix = "IVANAPI"

func main() {
	app := &cli.App{
		Name:  "Ivan's Portfolio API Server",
		Usage: "API Server to showcase features on portfolio",
		Commands: []*cli.Command{
			{
				Name:    "serve",
				Aliases: []string{"s"},
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "environment", Value: "development", EnvVars: []string{envPrefix + "_environment"}, Usage: "Environment as to what the server is running in"},
					&cli.StringFlag{Name: "api_addr", Value: ":8080", EnvVars: []string{envPrefix + "_API_ADDR", "API_ADDR"}, Usage: "Port for which server will run on"},
					&cli.StringFlag{Name: "log_level", Value: "DebugLevel", EnvVars: []string{envPrefix + "_LOG_LEVEL", "LOG_LEVEL"}, Usage: "Log level for logger"},
				},
				Usage: "Run server",
				Action: func(ctx *cli.Context) error {
					apiConfig := &api.Config{
						Address: ctx.String("api_addr"),
					}

					logger.NewLogger(ctx.String("environment"), ctx.String("log_level"))

					logger.L.Info().Msg("Starting API Server")
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
