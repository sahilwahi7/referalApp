package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/sahilwahi7/referalApp/connection"
	"github.com/sahilwahi7/referalApp/handler"
	"github.com/sahilwahi7/referalApp/repo"
)

var (
	once       sync.Once
	repository repo.Repo
	connect    connection.Connection
)

func getRepository() repo.Repo {
	once.Do(func() {
		repository = &repo.Concreterepo{}
	})
	return repository
}

func main() {
	h := &handler.Myhandler{
		Repo: getRepository(),
	}

	r := mux.NewRouter()

	r.HandleFunc("/viewJobs/{company}", h.ViewOpenJobs).Methods("GET")
	r.HandleFunc("/postJob", h.PostNewJob).Methods("POST")
	r.HandleFunc("/login", h.AuthServiceLogin).Methods("POST")
	r.HandleFunc("/signup", h.AuthServiceSignup).Methods("POST")
	r.HandleFunc("/updateProfile/{userName}", h.UpdateProfile).Methods("POST")

	// CORS middleware
	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},                      // Adjust as needed
		AllowedMethods: []string{"GET", "POST", "OPTIONS"}, // Adjust as needed
		AllowedHeaders: []string{"Content-Type"},           // Adjust as needed
		Debug:          true,
	})

	// Wrap the router with the CORS middleware
	handler := corsMiddleware.Handler(r)

	fmt.Printf("Connected to server, Welcome to referrals app")
	http.ListenAndServe(":8080", handler)
}
