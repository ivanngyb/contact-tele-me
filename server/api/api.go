package api

import (
	"context"
	"errors"
	"log"
	"net"
	"net/http"
	"server/cmd/server/telegram"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type API struct {
	Routes chi.Router
	Config *Config

	ctx    context.Context
	server *http.Server
}

type Config struct {
	Address string
}

func NewAPI(ctx context.Context, config *Config) *API {
	api := &API{
		Routes: chi.NewRouter(),
		Config: config,

		ctx: ctx,
	}

	api.Routes.Use(cors.New(
		cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedHeaders:   []string{"*"},
			AllowCredentials: true,
		}).Handler,
	)

	api.Routes.Route("/api", func(r chi.Router) {
		r.Get("/check", WithError(func(w http.ResponseWriter, r *http.Request) (int, error) {
			isAlive := telegram.TelegramCheck()

			if !isAlive {
				return http.StatusBadGateway, errors.New("telegram bot seems to be down")
			}

			return http.StatusOK, nil
		}))
		r.Post("/send", WithError(SendTeleMessage))
	})

	return api
}

func (api *API) Run(ctx context.Context) error {
	api.server = &http.Server{
		Addr:    api.Config.Address,
		Handler: api.Routes,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}

	go func() {
		<-ctx.Done()
		api.Close()
	}()

	return api.server.ListenAndServe()
}

func (api *API) Close() {
	ctx, cancel := context.WithTimeout(api.ctx, 5*time.Second)
	defer cancel()
	err := api.server.Shutdown(ctx)
	if err != nil {
		log.Fatal("failed to shutdown")
	}
}
