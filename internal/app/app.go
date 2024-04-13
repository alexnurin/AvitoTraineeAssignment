package app

import (
	"context"
	"fmt"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/api"
	"github.com/alexnurin/AvitoTraineeAssignment/internal/config"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

type Application struct {
	cfg    *config.Config
	db     *sqlx.DB
	router *gin.Engine
}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Start(ctx context.Context) error {
	if err := a.initConfig(); err != nil {
		return fmt.Errorf("can't init config: %w", err)
	}

	if err := a.initDatabaseConnection(); err != nil {
		return fmt.Errorf("can't init db connection: %w", err)
	}

	if err := a.initRouter(); err != nil {
		return fmt.Errorf("can't init router: %w", err)
	}
	return nil
}

func (a *Application) initConfig() error {
	var err error

	a.cfg, err = config.ParseConfig()
	if err != nil {
		return fmt.Errorf("failed to parse config: %w", err)
	}

	return nil
}

func (a *Application) initRouter() error {
	a.router = api.NewRouter()
	api.InitializeRoutes(a.router, a.db)
	err := a.router.Run(a.cfg.URL)
	if err != nil {
		return fmt.Errorf("failed to run router: %w", err)
	}

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

	err = dbConn.Ping()
	if err != nil {
		return err
	}

	a.db = dbConn
	log.Println("Database connection established")

	return nil
}
