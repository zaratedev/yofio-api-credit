package main

import (
    "log"
    "net/http"
	"os"

    "yofio-api/credit/handlers"

	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/aws/aws-lambda-go/lambda"
    "github.com/gorilla/mux"
	"github.com/joho/godotenv"
)


func main() {
	godotenv.Load()

    r := mux.NewRouter()

    r.HandleFunc("/credit-assigment", handlers.AssigmentInvesment).Methods("POST")

	if os.Getenv("MODE") == "dev" {
		log.Println("Run server in port 8000")
    	log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), r))

		return
	}

	adapter := gorillamux.New(r)

	lambda.Start(adapter.Proxy)
}
