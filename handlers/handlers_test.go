package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// Dummy logger
var l = log.New(ioutil.Discard, "", 0)

func TestHelloWorld(t *testing.T) {

	// HelloWorld Handler
	hw := &HelloWorldHandler{l}

	// Create a new ResponseRecorder
	rr := httptest.NewRecorder()
	// Create a dummy http.Request
	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create new handler for the Hello function and calls it
	handler := http.HandlerFunc(hw.Hello)
	handler.ServeHTTP(rr, req)

	// Get the handler response
	rs := rr.Result()

	// Check if we get 200
	if rs.StatusCode != http.StatusOK {
		t.Errorf("wrong handler status code: got %v want %v", rs.StatusCode, http.StatusOK)
	}

	expected := "Hello World!"
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	// Check if we get 'Hello World!'
	if string(body) != expected {
		t.Errorf("unexpected body response: got %v want %v", rr.Body.String(), expected)
	}
}

func TestGetCharacterByID(t *testing.T) {
	// HTTP client for requests
	client := &http.Client{}
	// HelloWorld Handler
	rm := &RickMortyHandler{l, client}

	// Create a new ResponseRecorder
	rr := httptest.NewRecorder()
	// Create a dummy http.Request
	req, err := http.NewRequest(http.MethodGet, "/rickmorty/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Need to create a router that we can pass the request through so that the vars will be added to the context
	router := mux.NewRouter()
	router.HandleFunc("/rickmorty/{id}", rm.GetCharacterByID)
	router.ServeHTTP(rr, req)

	// Get the handler response
	rs := rr.Result()

	// Check if we get 200
	if rs.StatusCode != http.StatusOK {
		t.Errorf("wrong handler status code: got %v want %v", rs.StatusCode, http.StatusOK)
	}

	expected := "Rick Sanchez"
	body, err := ioutil.ReadAll(rs.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer rs.Body.Close()

	if !strings.Contains(string(body), expected) {
		t.Errorf("body response doesn't contain expected string: got %v want %v", rr.Body.String(), expected)
	}
}
