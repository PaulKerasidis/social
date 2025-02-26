package main

import (
	"log"

	"github.com/PaulKerasidis/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADDR", ":8989"),
	}

	app := &application{
		config: cfg,
	}

	

	mux := app.mount()

	log.Fatal(app.run(mux))
}