package main

import (
	"database/sql"
	_ "net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {
	InitializeDB()

	// Create a new router
	r := gin.Default()

	r.GET("/", TestAPI)
	r.GET("/api", TestAPI)

	r.GET("/api/disease", GetDisease)
	r.GET("/api/disease/:id", GetDiseaseById)
	// Menerima string "disease_name" dan file "dna_sequence_file"
	r.POST("/api/disease", InsertDisease)
	r.DELETE("/api/disease/:id", DeleteDiseaseById)

	r.GET("/api/prediction", GetPrediction)
	r.GET("/api/prediction/:id", GetPredictionById)
	// Menerima string "patient_name", file "dna_sequence_file", dan "disease_id"	
	r.POST("/api/prediction", InsertPrediction)
	r.DELETE("/api/prediction/:id", DeletePredictionById)

	r.Run(get_port())
}

func get_port() string {
    port := ":8080"
    if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
        port = ":" + val
    }
    return port
}
