package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gorilla/mux"
)

var CharactersEndpoint = "https://rickandmortyapi.com/api/character/"

type RickMortyHandler struct {
	logger *log.Logger
	client *http.Client
}

type Character struct {
	ID       int       `json:"id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Status   string    `json:"status,omitempty"`
	Species  string    `json:"species,omitempty"`
	Type     string    `json:"type,omitempty"`
	Location *Location `json:"location,omitempty"`
	Origin   *Origin   `json:"origin,omitempty"`
	Error    string    `json:"error,omitempty"`
}

type Location struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type Origin struct {
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

// NewRickMortyHandler returns a RickMortyHandler struct with a logger.
func NewRickMortyHandler(l *log.Logger, c *http.Client) *RickMortyHandler {
	return &RickMortyHandler{l, c}
}

// GetCharacterByID returns information about a Rick and Morty character using a custom struct
func (rm *RickMortyHandler) GetCharacterByID(w http.ResponseWriter, r *http.Request) {
	// Get the route variables
	vars := mux.Vars(r)
	id := vars["id"]
	// Complete the endpoint with the character id
	url := CharactersEndpoint + id

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		rm.logger.Printf("msg: error getting character, err: %v\n", err)
		serverError(w, rm.logger, err)
		return
	}
	// Send the HTTP request
	resp, err := rm.client.Do(req)
	if err != nil {
		rm.logger.Printf("msg: error sending http request, err: %v\n", err)
		serverError(w, rm.logger, err)
		return
	}
	// Close the body at the end of execution
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		rm.logger.Printf("msg: error while reading the body, err: %v\n", err)
		serverError(w, rm.logger, err)
		return
	}

	// Empty struct to hold the response data
	char := &Character{}

	err = json.Unmarshal(body, char)
	if err != nil {
		rm.logger.Printf("msg: error while unmarshaling character data, err: %v\n", err)
		serverError(w, rm.logger, err)
		return
	}

	jsonChar, err := json.Marshal(char)
	if err != nil {
		rm.logger.Printf("msg: error while marshaling character data, err: %v\n", err)
		serverError(w, rm.logger, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "%v\n", string(jsonChar))
}

// Helper function to get a trace of an error and reply with a 500:
func serverError(w http.ResponseWriter, errorLog *log.Logger, err error) {
	//Debug trace
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
