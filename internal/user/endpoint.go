package user

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller func(w http.ResponseWriter, r *http.Request)
type Endpoints struct {
	Create Controller
	Get    Controller
	GetAll Controller
	Update Controller
	Delete Controller
}
type CreateRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		Create: makeCreateEndpoint(s),
		Get:    makeGetEndpoint(s),
		GetAll: makeGetAllEndpoint(s),
		Update: makeUpdateEndpoint(s),
		Delete: makeDeleteEndpoint(s),
	}
}

func makeCreateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {

		// fmt.Println("Create user")

		var req CreateRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			fmt.Println("Error decoding request", err)
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request: " + err.Error()})
			return
		}

		if req.FirstName == "" || req.LastName == "" || req.Email == "" || req.Phone == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid request: missing required fields"})
			return
		}

		User, err := s.Create(req.FirstName, req.LastName, req.Email, req.Phone)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(ErrorResponse{Error: "Error creating user: " + err.Error()})
			return
		}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(User)
	}
}

func makeGetEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Get user")

		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeGetAllEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("GetAll user")

		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}

func makeUpdateEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Update user")

		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
func makeDeleteEndpoint(s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Delete user")

		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	}
}
