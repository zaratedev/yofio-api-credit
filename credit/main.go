package main

import (
    "context"
    "log"
    "net/http"
    "os"

    "yofio-api/credit/handlers"

    "github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

func main() {
    // Load env
    godotenv.Load()

    // Set mongo db client
    clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(context.Background())

    collection = client.Database("yofio-credit").Collection("assigment_credit")

    // Router
    r := mux.NewRouter()

    r.HandleFunc("/credit-assigment", func(w http.ResponseWriter, r *http.Request) {
        handlers.AssigmentInvestment(w, r, collection)
    }).Methods("POST")

    r.HandleFunc("/statistics", func(w http.ResponseWriter, r *http.Request) {
        handlers.StatisticsHandler(w, r, collection)
    }).Methods("POST")
	
    if os.Getenv("MODE") == "dev" {
        log.Println("Run server in port 8000")
        log.Fatal(http.ListenAndServe(os.Getenv("SERVER_PORT"), r))
        return
    }

    adapter := gorillamux.New(r)
    lambda.Start(adapter.Proxy)
}
