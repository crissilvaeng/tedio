package core

import (
	"fmt"
	"log"
	"os"

	"github.com/crissilvaeng/tedio/internal/storage"
)

type Server struct {
	Auth    *Auth
	Routes  *Routes
	Logger  *log.Logger
	Address string
}

type Config struct {
	ApiKey    string
	ApiSecret string
	Port      string
}

func NewServer(config Config) (*Server, error) {
	logger := log.New(os.Stdout, "api: ", log.LstdFlags)
	server := &Server{
		Auth: &Auth{
			apikey: config.ApiKey,
			secret: config.ApiSecret,
		},
		Routes: &Routes{
			repository: storage.NewInMemoryStorage(),
		},
		Logger:  logger,
		Address: fmt.Sprintf("http://localhost:%s", config.Port),
	}
	return server, nil
}
