package app

import (
	"errors"
	"fmt"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/api"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/config"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
	"time"
)

func (a *Application) initConfig() error {
	var err error

	a.cfg, err = config.ParseConfig()
	if err != nil {
		return fmt.Errorf("не удалось загрузить конфигурацию: %w", err)
	}

	return nil
}

func (a *Application) initRouter() error {
	a.router = api.NewRouter()
	api.InitializeRoutes(a.router, a.db)
	//err := a.router.Run(a.cfg.URL)
	//if err != nil {
	//	return fmt.Errorf("ошибка при запуске роутера: %w", err)
	//}

	return nil
}

func (a *Application) initServer() error {
	a.httpServer = &http.Server{
		Addr:    a.cfg.HTTPServer.URL,
		Handler: a.router,
	}
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %w", err)
		}
	}()
	return nil
}

func (a *Application) initDatabaseConnection() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		a.cfg.DB.Host,
		a.cfg.DB.Port,
		a.cfg.DB.User,
		a.cfg.DB.Password,
		a.cfg.DB.Name)

	dbConn, err := sqlx.Open("postgres", psqlInfo)
	if err != nil {
		return err
	}

	dbConn.SetMaxOpenConns(25)           // Максимальное количество открытых соединений
	dbConn.SetMaxIdleConns(10)           // Максимальное количество простаивающих соединений
	dbConn.SetConnMaxLifetime(time.Hour) // Максимальное время жизни соединения

	err = dbConn.Ping()
	if err != nil {
		return err
	}

	a.db = dbConn
	log.Println("Database connection established")

	return nil
}
