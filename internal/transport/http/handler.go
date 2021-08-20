package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Handler Stores pointer to our comments services
type Handler struct {
	Router *mux.Router
}

// NewHandler - Returns pointer to a Handler
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - set up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am Alive !!!!")
	})
}
