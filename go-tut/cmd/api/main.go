package main

import (
	"fmt"
	"log"

	"github.com/samson4/go-tut/internal/env"
)

func main() {
	fmt.Println("PORT: ", env.GetString("PORT", ""))
	cfg := config{
		addr: env.GetString("PORT", ":8080"),
	}
	app := &application{
		config: cfg,
	}
	mux := app.mount()
	log.Fatal(app.run(mux))
}
