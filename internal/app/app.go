package app

import (
	"context"
	"fmt"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type Application struct {
	cfg        *config.Config
	db         *sqlx.DB
	router     *gin.Engine
	httpServer *http.Server
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Start() error {
	if err := a.initConfig(); err != nil {
		return fmt.Errorf("не удалось загрузить конфигурацию: %v", err)
	}

	if err := a.initDatabaseConnection(); err != nil {
		return fmt.Errorf("не удалось подключиться к базе данных: %v", err)
	}

	if err := a.initRouter(); err != nil {
		return fmt.Errorf("не удалось запустить роутер: %v", err)
	}

	if err := a.initServer(); err != nil {
		return fmt.Errorf("не удалось запустить сервер: %v", err)
	}
	return nil
}

func (a *Application) Wait(ctx context.Context) error {
	<-ctx.Done()

	if err := a.db.Close(); err != nil {
		log.Printf("не удалось закрыть соединение с базой данных: %v", err)
	}

	if err := a.httpServer.Shutdown(context.Background()); err != nil {
		log.Printf("не удалось корректно остановить HTTP сервер: %v", err)
	}

	log.Println("Все системы корректно завершили работу.")
	return nil
}
