package main

import (
	"fmt"
	"os"

	"go.uber.org/fx"
)

/*
API
AuthService
Database	Logger	Config
*/

/* podia ser com interface
type DatabaseInterface interface {
	GetUser()
}

func NewDatabas2() *DatabaseInterface {
	return &Database{}
}
*/

type Database struct{}

// Construtor
func NewDatabase() *Database {
	return &Database{}
}

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

type Config struct {
	Name string
}

func NewConfig() *Config {
	return &Config{
		Name: os.Getenv("NAME"),
	}
}

// os campos são as dependências
type AuthService struct {
	Database *Database
	Logger   *Logger
	Config   *Config
}

func NewAuthService(database *Database, logger *Logger, config *Config) *AuthService {
	return &AuthService{
		Database: database,
		Logger:   logger,
		Config:   config,
	}
}

type API struct {
	AuthService *AuthService
}

func NewAPI(authService *AuthService) *API {
	return &API{
		AuthService: authService,
	}
}

func main() {
	/* padrao
	database := NewDatabase()
	logger := NewLogger()
	config := NewConfig()
	authService := NewAuthService(database, logger, config)
	_ = NewAPI(authService) // alto nível da dependência
	*/

	// colocar na ordem de prioridades
	fx.New(
		fx.Provide(NewDatabase, NewLogger, NewConfig, NewAuthService, NewAPI),
		fx.Invoke(func(api *API) {
			fmt.Println("todas as dependencias foram inicilizadas")
		}),
	).Run()
}
