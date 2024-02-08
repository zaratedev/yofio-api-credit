package handlers

import (
    "context"
    "encoding/json"
    "net/http"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

type Statistic struct {
    Total               int     `json:"total"`
    Successful          int     `json:"successful"`
    Unsuccessful        int     `json:"unsuccessful"`
    AverageSuccessful   float64 `json:"average_successful"`
    AverageUnsuccessful float64 `json:"average_unsuccessful"`
}

func StatisticsHandler(w http.ResponseWriter, r *http.Request, collection *mongo.Collection) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    total, successful, unsuccessful, avgSuccessful, avgUnsuccessful, err := calculateStatistics(collection)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    stats := Statistic{
        Total:              total,
        Successful:         successful,
        Unsuccessful:       unsuccessful,
        AverageSuccessful: avgSuccessful,
        AverageUnsuccessful: avgUnsuccessful,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(stats)
}

func calculateStatistics(collection *mongo.Collection) (int, int, int, float64, float64, error) {
    total, err := collection.CountDocuments(context.Background(), bson.D{})
    if err != nil {
        return 0, 0, 0, 0, 0, err
    }

    // Count Sucessfull assigment
    successful, err := collection.CountDocuments(context.Background(), bson.D{{"successful", true}})
    if err != nil {
        return 0, 0, 0, 0, 0, err
    }

    // Count unsucessfull assigment
    unsuccessful := total - successful

    // Avg assigment sucessfull
    pipeline := bson.A{
        bson.D{{"$match", bson.D{{"successful", true}}}},
        bson.D{
            {"$group",
                bson.D{
                    {"_id", nil},
                    {"average", bson.D{{"$avg", "$investment"}}},
                },
            },
        },
    }

    cursor, err := collection.Aggregate(context.Background(), pipeline)
    if err != nil {
        return 0, 0, 0, 0, 0, err
    }
    defer cursor.Close(context.Background())

    var result struct {
        Average float64 `bson:"average"`
    }
    if cursor.Next(context.Background()) {
        if err := cursor.Decode(&result); err != nil {
            return 0, 0, 0, 0, 0, err
        }
    }

    avgSuccessful := result.Average

    // Avg assigment unsucessfull
    avgUnsuccessful := 0.0
    if unsuccessful > 0 {
        pipeline = bson.A{
            bson.D{{"$match", bson.D{{"successful", false}}}},
            bson.D{
                {"$group",
                    bson.D{
                        {"_id", nil},
                        {"average", bson.D{{"$avg", "$investment"}}},
                    },
                },
            },
        }

        cursor, err = collection.Aggregate(context.Background(), pipeline)
        if err != nil {
            return 0, 0, 0, 0, 0, err
        }
        defer cursor.Close(context.Background())

        if cursor.Next(context.Background()) {
            if err := cursor.Decode(&result); err != nil {
                return 0, 0, 0, 0, 0, err
            }
        }
        avgUnsuccessful = result.Average
    }

    return int(total), int(successful), int(unsuccessful), avgSuccessful, avgUnsuccessful, nil
}
