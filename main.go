package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Aeross/wbootcamp/config"
)

func main() {
	// Flag to change the server address if required
	addr := flag.String("addr", ":4000", "HTTP network address")
	// Parse the command-line flags, in this case 'addr'
	flag.Parse()

	// Logger
	l := log.New(os.Stdout, "WBOOTCAMP\t", log.Ldate|log.Ltime)

	// Init Config
	config := config.New(l)

	// Init serve mux routes
	sm := config.InitRoutes()

	// HTTP Server configuration
	s := &http.Server{
		Addr:         *addr,             //set the bind address
		Handler:      sm,                //set the default handler
		ReadTimeout:  5 * time.Second,   //max time to read request from the client
		WriteTimeout: 10 * time.Second,  //max time to write request from the client
		IdleTimeout:  120 * time.Second, //max time for connections using TCP Keep-Alive
		ErrorLog:     l,                 //set the logger for the server
	}

	// Startup messages
	l.Printf("Starting server on port %s", *addr)
	l.Printf("Go to server http://localhost%s", *addr)
	l.Printf("Get a hello world response: http://localhost%s/hello", *addr)
	l.Printf("Get information about a Rick and Morty character. i.e: http://localhost%s/rickmorty/1", *addr)

	// Anon function to start the HTTP Server.
	go func() {
		// It returns an error in case of any failure.
		err := s.ListenAndServe()
		if err != nil {
			l.Fatalf("msg: error initializing HTTP Server, err: %v\n", err)
		}
	}()

	// Listen for os signals to terminate the server.
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	// If the server was interrupted, print the interruption signal that stopped it.
	sig := <-sigChan
	l.Println("msg: received terminate, graceful shutdown. signal:", sig)

	// Wait until the server has finished processing pending jobs, then terminate. Or, if the context time threshold is met, terminate.
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
