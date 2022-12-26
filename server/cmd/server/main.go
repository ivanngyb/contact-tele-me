package main

import (
	"log"
	"os"
	"server/api"
	"server/cmd/server/telegram"
	"server/logger"

	"github.com/urfave/cli/v2"
)

const envPrefix = "CONTACT"

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
					&cli.StringFlag{Name: "api_addr", Value: ":8084", EnvVars: []string{envPrefix + "_API_ADDR", "API_ADDR"}, Usage: "Port for which server will run on"},
					&cli.StringFlag{Name: "log_level", Value: "DebugLevel", EnvVars: []string{envPrefix + "_LOG_LEVEL", "LOG_LEVEL"}, Usage: "Log level for logger"},
					&cli.StringFlag{Name: "bot_token", Value: "", EnvVars: []string{envPrefix + "_BOT_TOKEN"}, Usage: "Telegram Bot API Token"},
					&cli.Int64Flag{Name: "bot_chat_id", Value: 0, EnvVars: []string{envPrefix + "_BOT_CHAT_ID"}, Usage: "Telegram Bot Chat ID"},
				},
				Usage: "Run server",
				Action: func(ctx *cli.Context) error {
					apiConfig := &api.Config{
						Address: ctx.String("api_addr"),
					}

					logger.NewLogger(ctx.String("environment"), ctx.String("log_level"))

					err := telegram.NewBot(ctx.String("bot_token"), ctx.Int64("bot_chat_id"))
					if err != nil {
						logger.L.Fatal().Msg("Failed to start telegram bot")
					}

					logger.L.Info().Msg("Starting API Server")
					api := api.NewAPI(ctx.Context, apiConfig)

					err = api.Run(ctx.Context)
					if err != nil {
						logger.L.Fatal().Msg("fail to start api")
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
