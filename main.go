package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/AhmedOthman94/training-exercises/internal/app"
	"github.com/AhmedOthman94/training-exercises/internal/routes"
)

func main() {
	// Add port flag to command line and set default to 8080
	var port int
	flag.IntVar(&port, "port", 8080, "go backend server port")
	flag.Parse()

	// Create new application
	app, err := app.NewApplication()
	if err != nil {
		panic(err)
	}

	// Setuo routes
	r := routes.SetupRoutes(app)

	// Create new server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      r,
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.Logger.Printf("We are running out app on port: %d\n", port)

	// Start server
	if err := server.ListenAndServe(); err != nil {
		app.Logger.Fatal(err)
	}

}
