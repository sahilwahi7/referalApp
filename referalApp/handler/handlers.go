package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/sahilwahi7/referalApp/connection"
	"github.com/sahilwahi7/referalApp/repo"
)

type Myhandler struct {
	Repo    repo.Repo
	Connect connection.Connection
}

// main handler interface that exposes all endpoints
type Handler interface {
	AuthServiceLogin(w http.ResponseWriter, r *http.Request)  //completed
	AuthServiceSignup(w http.ResponseWriter, r *http.Request) //completed
	PostNewJob(w http.ResponseWriter, r *http.Request)
	// ApplyJob(w http.ResponseWriter, r *http.Request)
	ViewOpenJobs(w http.ResponseWriter, r *http.Request)
	UpdateProfile(w http.ResponseWriter, r *http.Request)
}

func (h *Myhandler) UpdateProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userName := vars["userName"]
	var j repo.UpdatedUser
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		fmt.Printf("Invalid job body...")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	updatedProfile := h.Repo.UpdateUserProfile(userName, j.LinkedProfile, j.CurrentCompany, j.Description, j.Title)
	response, err := json.Marshal(updatedProfile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("There is some error updating your profile..please try again")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func (h *Myhandler) ViewOpenJobs(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	company := vars["company"]
	jobs := h.Repo.FindOpenJobs(company)
	reponse, err := json.Marshal(jobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error please contact support")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(reponse)

}

func (h *Myhandler) PostNewJob(w http.ResponseWriter, r *http.Request) {
	var j repo.Data
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		fmt.Printf("Invalid job body...")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	//TO-DO here we will fetch the parameter from UI, for now passing default as name
	jobs := h.Repo.FindOpenJobs(j.Compnay)
	jobs = append(jobs, j)
	//h.Repo.StoreJob(j)
	//need to store this j job in jobs.Models also for viewall jobs

	response, err := json.Marshal(jobs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error please contact support")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func (h *Myhandler) AuthServiceLogin(w http.ResponseWriter, r *http.Request) {
	var j repo.Authentication
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		fmt.Printf("Invalid Authentication body....")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	resp := h.Repo.Login(j.UserName, j.Password)
	response, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error please contact support")

	}
	// here we will implement the logic of signup or login
	// if username exists in DB then authenticate else ask for signup
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func (h *Myhandler) AuthServiceSignup(w http.ResponseWriter, r *http.Request) {
	var j repo.User
	if err := json.NewDecoder(r.Body).Decode(&j); err != nil {
		fmt.Printf("Invalid Authentication body....")
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	resp := h.Repo.Signup(j.Name, j.UserName, j.Password, j.ID, j.IsRefree)
	response, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Printf("Error please contact support")
		return
	}
	// here we will implement the logic of signup or login
	// if username exists in DB then authenticate else ask for signup
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

// func (h *Myhandler) ViewUsers(w http.ResponseWriter, r *http.Request) {
// 	users := h.Repo.ViewAllUsers()
// 	reponse, err := json.Marshal(users)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		fmt.Printf("Error please contact support")
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(reponse)
// }
