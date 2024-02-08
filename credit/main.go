package main

import (
    "log"
    "net/http"
	"os"

    "yofio-api/credit/handlers"

	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/gorilla/mux"
)


func main() {
    r := mux.NewRouter()

    r.HandleFunc("/credit-assigment", handlers.AssigmentInvesment).Methods("POST")

	if os.Getenv("MODE") == "dev" {
		log.Println("Run server in port 8000")
    	log.Fatal(http.ListenAndServe(":8000", r))

		return
	}

	adapter := gorillamux.New(r)

	lambda.Start(adapter.Proxy)
}
