package main

import (
    "log"
    "net/http"

    "github.com/gorilla/mux"
    "yofio-api/handlers"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/credit-assigment", handlers.AssigmentInvesment).Methods("POST")

    log.Println("Run server in port 8000")
    log.Fatal(http.ListenAndServe(":8000", r))
}
