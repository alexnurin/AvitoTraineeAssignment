package main

import (
	"context"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/app"
	"runtime"
	"time"

	"log"
	"os/signal"
	"syscall"
)

func main() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGHUP,
		syscall.SIGQUIT,
	)
	defer stop()

	numCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPU)
	log.Printf("GOMAXPROCS set to %d", numCPU)

	application := app.NewApplication()
	if err := application.Start(); err != nil {
		log.Fatalf("не удалось запустить приложение: %v", err)
	}

	if err := application.Wait(ctx); err != nil {
		log.Printf("приложение завершило работу с ошибкой: %v", err)
	} else {
		log.Println("Приложение завершило работу без ошибок")
	}

	time.Sleep(1 * time.Second)
}
