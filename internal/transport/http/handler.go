package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/phanimullapudi/go-rest-api-course/internal/comment"
)

// Handler Stores pointer to our comments services
type Handler struct {
	Router  *mux.Router
	Service *comment.Service
}

// Response - a object to store responses from API
type Response struct {
	Message string
	Error   string
}

// NewHandler - Returns pointer to a Handler
func NewHandler(service *comment.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

// SetupRoutes - set up all the routes for our application
func (h *Handler) SetupRoutes() {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/comment", h.GetAllComments).Methods("GET")
	h.Router.HandleFunc("/api/comment", h.PostComment).Methods("POST")
	h.Router.HandleFunc("/api/comment/{id}", h.GetComment).Methods("GET")
	h.Router.HandleFunc("/api/comment/{id}", h.UpdateComment).Methods("PUT")
	h.Router.HandleFunc("/api/comment/{id}", h.DeleteComment).Methods("DELETE")

	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "Iam ALIVE !"}); err != nil {
			panic(err)
		}
	})
}

// GetComment - retrieve comment from the comment service by ID
func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
		return

	}

	comment, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving comment by ID", err)
		return
	}

	if err := json.NewEncoder(w).Encode(comment); err != nil {
		panic(err)
	}
}

// GetAllComments - retrieves all comments from the comment service
func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := h.Service.GetAllComments()
	if err != nil {
		sendErrorResponse(w, "Failed to get all comments", err)
		return
	}
	if err := sendOkResponse(w, comments); err != nil {
		panic(err)
	}
}

// PostComment - Post a new comment
func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to DECODE json body", err)
		return
	}

	comment, err := h.Service.PostComment(comment)
	if err != nil {
		sendErrorResponse(w, "Failed to post new comment", err)
		return
	}
	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// UpdateComment - update a comment by ID
func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var comment comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		sendErrorResponse(w, "Failed to DECODE json body", err)
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse unint from ID", err)
		return
	}

	comment, err = h.Service.UpdateComment(uint(commentID), comment)
	if err != nil {
		sendErrorResponse(w, "Failed to update comment", err)
		return
	}

	if err := sendOkResponse(w, comment); err != nil {
		panic(err)
	}
}

// DeleteComment - deletes a comment by ID
func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	commentID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse unint from ID", err)
		return
	}

	err = h.Service.DeleteComment(uint(commentID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete the comment by ID", err)
		return
	}

	if err = sendOkResponse(w, Response{Message: "Successfully Deleted"}); err != nil {
		panic(err)
	}
}

func sendOkResponse(w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
}

func sendErrorResponse(w http.ResponseWriter, messsage string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: messsage, Error: err.Error()}); err != nil {
		panic(err)
	}
}
