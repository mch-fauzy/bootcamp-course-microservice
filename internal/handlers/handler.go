package handlers

import (
	"encoding/json"
	"net/http"

	"bootcamp-course-microservice/internal/models"
	"bootcamp-course-microservice/internal/services"
)

type Handler struct {
	Service services.Service
}

func ProvideHandler(service services.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	// Define the required struct for the request body
	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}

	// Decode the request body into the req struct
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate required fields
	if req.Title == "" || req.Content == "" {
		http.Error(w, "title and content fields are required", http.StatusBadRequest)
		return
	}

	//Might need to check if ID is exist or not (to avoid duplicate)

	// Create the product model from the request data
	course := &models.Course{
		UserID:  r.Context().Value("user_id").(string),
		Title:   req.Title,
		Content: req.Content,
	}

	err := h.Service.CreateCourse(course)
	if err != nil {
		http.Error(w, "Failed to create course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	response := map[string]interface{}{
		"message": "Course successfully created",
	}
	json.NewEncoder(w).Encode(response)
}

func (h *Handler) ReadCourseByUserID(w http.ResponseWriter, r *http.Request) {
	// Retrieve the user ID from the context
	UserID := r.Context().Value("user_id").(string)

	course, err := h.Service.ReadCourseByUserID(UserID)
	if err != nil {
		http.Error(w, "Failed to fetch course", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(course)
}
