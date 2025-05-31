package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"vertical-slice-arch/internal/app"
)

func main() {
	ctx := context.Background()
	exampleApp := app.NewApplication()

	if err := exampleApp.Init(ctx); err != nil {
		log.Fatal(err)
		return
	}

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, os.Kill)
	defer cancel()

	if err := exampleApp.Run(ctx); err != nil {
		return
	}
}
