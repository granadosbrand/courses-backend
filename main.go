package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()


	router.HandleFunc("/users", getUsers)
	router.HandleFunc("/courses", getCourses)

	srv := &http.Server{
		Handler: router,
		Addr: "127.0.0.1:8000",
	}

	log.Fatal(srv.ListenAndServe())
} 

func getUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	
}
func getCourses(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]bool{"ok": true})
}
