package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutsHandler struct{}

func NewWorkoutsHandler() *WorkoutsHandler {
	return &WorkoutsHandler{}
}

func (wh *WorkoutsHandler) HandleGetWorkoutByID(w http.ResponseWriter, r *http.Request) {
	paramWorkoutID := chi.URLParam(r, "id")
	if paramWorkoutID == "" {
		http.NotFound(w, r)
		return
	}

	workoutID, err := strconv.ParseInt(paramWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Workout ID: %d\n", workoutID)
}

func (wh *WorkoutsHandler) HandleCreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Workout created\n")
}
