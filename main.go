package main

import (
	"gpu-container-manager/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/containers", handlers.CreateContainer).Methods("POST")
    r.HandleFunc("/containers", handlers.UpdateContainer).Methods("PUT")
    r.HandleFunc("/containers", handlers.DeleteContainer).Methods("DELETE")

    http.ListenAndServe(":8080", r)
}
