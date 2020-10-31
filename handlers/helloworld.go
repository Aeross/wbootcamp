package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// Define the HelloWorld struct handler
type HelloWorldHandler struct {
	logger *log.Logger
}

// NewHelloWorldHandler returns a HelloWorldHandler struct with a logger.
func NewHelloWorldHandler(l *log.Logger) *HelloWorldHandler {
	return &HelloWorldHandler{l}
}

// Handler function to return Hello World
func (hw *HelloWorldHandler) Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
