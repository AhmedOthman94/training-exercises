package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AhmedOthman94/training-exercises/internal/api"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutsHandler
}

func NewApplication() (*Application, error) {
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	// store

	// Create new workout handler
	workoutHandler := api.NewWorkoutsHandler()

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
	}

	return app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status: OK\n")
}
