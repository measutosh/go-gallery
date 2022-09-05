// Goals in a flow

// => just to make a router that will interact with the frontend(made in routes.go)
// => a receiever will catch the responses that is built here
package main

import (
	"log"
	"fmt"
	"net/http"
)

const webPort = "80"

// declaring type for the receiever
type Config struct {}


func main() {
	app := Config{}

	log.Printf("Starting broker service at port %s", webPort)

	// set up the http server
	srv := &http.Server {
		Addr : fmt.Sprintf(":%s", webPort),
		// call routes.go
		Handler: app.routes(),
	}

	// start the server
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}