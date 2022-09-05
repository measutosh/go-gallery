// this contains all the routes for the application
package main

import (
	// "fmt"clear
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

// here (app *Config) is the receiever that comes from main.go that can help fowarding any kind of data
func (app *Config) routes() http.Handler {
	mux := chi.NewRouter()

	// mention who is allowed to connect
	mux.Use(cors.Handler(cors.Options{
		AllowedOrigins:     []string{"http://*", "https://*"},
		AllowedMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:     []string{"Link"},
		AllowCredentials:   true,
		MaxAge:             300,
	}))

	// for a aliveness check-up, this should respond something to prove it's active
	mux.Use(middleware.Heartbeat("/ping"))
    mux.Post("/", app.Broker)

	
	return mux
}
