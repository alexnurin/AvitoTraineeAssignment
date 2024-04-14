package main

import (
	"context"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/app"

	"log"
	"os/signal"
	"syscall"
)

func main() {
	_, stop := signal.NotifyContext( // TODO cancel
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	defer stop()

	application := app.NewApplication()

	if err := application.Start(); err != nil {
		log.Fatalf("не удалось запустить приложение: %v", err)
	}
	log.Println("Приложение завершило работу без ошибок")
}
