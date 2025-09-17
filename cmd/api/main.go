package main

import (
	"log"
	"os"

	"github.com/dslcosta1/social/internal/env"
)

func main() {
	cfg := config{
		addr: env.GetString("ADRR", ":8080"),
	}

	app := &application{
		config: cfg,
	}

	os.LookupEnv("PATH")

	mux := app.mount()

	log.Fatal(app.run(mux))
}
