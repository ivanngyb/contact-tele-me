package api

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
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

	api.Routes.Route("/api", func(r chi.Router) {
		r.Get("/check", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("OK"))
		})
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
