package config

import (
	"log"
	"net/http"

	"github.com/Aeross/wbootcamp/handlers"
	"github.com/gorilla/mux"
)

type Config struct {
	logger *log.Logger
}

func New(l *log.Logger) *Config {
	return &Config{l}
}

func (c *Config) InitRoutes() *mux.Router {
	// HTTP client for requests
	client := &http.Client{}
	// Handlers
	hw := handlers.NewHelloWorldHandler(c.logger)
	rm := handlers.NewRickMortyHandler(c.logger, client)

	// Gorilla serve mux
	sm := mux.NewRouter()

	// Gorilla routes
	// Get
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/hello", hw.Hello)
	getRouter.HandleFunc("/rickmorty/{id}", rm.GetCharacterByID)

	// I could have checked the character id with a regex to allow integers only,
	// but the Rick and Morty API already returns an error response with nice messages
	// therefore I used this error response in the Character struct as an additional Error property
	//
	// I could have done something like:
	// getRouter.HandleFunc("/rickmorty/{id:[0-9]+$}", rm.GetCharacterByID)

	return sm
}
