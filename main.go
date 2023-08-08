package main

import (
	"bootcamp-course-microservice/infras"
	"bootcamp-course-microservice/internal/repository"
	"bootcamp-course-microservice/internal/services"
	"bootcamp-course-microservice/transport/middleware"
	"bootcamp-course-microservice/transport/routes"
	"fmt"
	"net/http"
)

func main() {
	// Create a new database connection
	db := infras.ProvideConn()

	// Initialize the repository with the database connection
	repo := repository.ProvideRepo(&db)

	// Initialize the service with the repository
	svc := services.ProvideService(repo)

	// Initialize the authentication middleware
	secretKey := []byte("your-secret-key")
	auth := middleware.ProvideAuthentication(&db, secretKey)

	// Initialize the router with the service and authentication
	r := routes.ProvideRouter(svc, auth)

	fmt.Println("Starting server on :8081")
	err := http.ListenAndServe(":8081", r.SetupRoutes())
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
}
