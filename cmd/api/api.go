package main

import (
	"log"
	"net/http"
	"time"
)

type application struct {

	config config

}

type config struct {
	addr string

}

func (app *application) mount() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/health", app.healthCheckHandler)

	return mux
}

func (app *application) run(mux *http.ServeMux) error {

	srv := &http.Server{
		Addr: app.config.addr,
		Handler: mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout: time.Second * 10,
		IdleTimeout: time.Second * 60,
	}
	log.Printf("Starting server on http://localhost%s", srv.Addr)

	return srv.ListenAndServe()
}