package main

import (
	"context"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/app"

	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext( // TODO cancel
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	defer stop()

	application := app.NewApplication()

	if err := application.Start(ctx); err != nil {
		log.Fatalf("failed to run app: %v", err)
	}
	<-ctx.Done()
	log.Println("All systems closed without errors")

}
