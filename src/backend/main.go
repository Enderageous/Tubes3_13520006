package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	_ "net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	cfg := mysql.Config{
		User:   os.Getenv("DB_USER"),
		Passwd: os.Getenv("DB_PASSWORD"),
		Net:    os.Getenv("DB_NET"),
		Addr:   os.Getenv("DB_HOST"),
		DBName: os.Getenv("DB_NAME"),

		AllowNativePasswords: true,
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	checkError(err)
	pingErr := db.Ping()
	checkError(pingErr)
	fmt.Println("Connected to database")

	// Query the database
	// diseases, err := diseaseList()
	// checkError(err)
	// fmt.Println(diseases)

	// Create a new router
	r := gin.Default()
	r.GET("/api", func(c *gin.Context) {
		c.String(200, "200 OK")
	})

	r.GET("/api/disease", func(c *gin.Context) {
		// Get all diseases from the database.
		rows, err := db.Query("SELECT id, disease_name, dna_sequence FROM disease")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		// Iterate through the diseases and create a JSON array.
		var diseases []Disease
		for rows.Next() {
			var disease Disease
			if err := rows.Scan(&disease.ID, &disease.DiseaseName, &disease.DNASequence); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			diseases = append(diseases, disease)
		}
		// Return the diseases.
		c.JSON(200, gin.H{"diseases": diseases})
	})

	r.GET("/api/disease/:id", func(c *gin.Context) {
		// Get the id parameter from the URL.
		id := c.Params.ByName("id")
		// Get the disease from the database.
		var disease Disease
		err := db.QueryRow("SELECT id, disease_name, dna_sequence FROM disease WHERE id = ?", id).Scan(&disease.ID, &disease.DiseaseName, &disease.DNASequence)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Return the disease.
		c.JSON(200, disease)
	})

	r.GET("/api/prediction", func(c *gin.Context) {
		// Get all predictions from the database.
		rows, err := db.Query("SELECT id, date, patient_name, disease_name, result, accuracy FROM prediction_view")
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		defer rows.Close()
		// Iterate through the predictions and create a JSON array.
		var predictions []Prediction
		for rows.Next() {
			var prediction Prediction
			if err := rows.Scan(&prediction.ID, &prediction.Date, &prediction.PatientName, &prediction.DiseasePrediction, &prediction.Result, &prediction.Accuracy); err != nil {
				c.JSON(500, gin.H{"error": err.Error()})
				return
			}
			predictions = append(predictions, prediction)
		}
		// Return the predictions.
		c.JSON(200, gin.H{"predictions": predictions})
	})

	r.GET("/api/prediction/:id", func(c *gin.Context) {
		// Get the id parameter from the URL.
		id := c.Params.ByName("id")
		// Get the prediction from the database.
		var prediction Prediction
		err := db.QueryRow("SELECT id, date, patient_name, disease_name, result, accuracy FROM prediction_view WHERE id = ?", id).Scan(&prediction.ID, &prediction.Date, &prediction.PatientName, &prediction.DiseasePrediction, &prediction.Result, &prediction.Accuracy)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Return the prediction.
		c.JSON(200, prediction)
	})

	r.POST("/api/disease", func(c *gin.Context) {
		// Capture the disease from the request body.
		var newDisease Disease
		if err := c.Bind(&newDisease); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		file, err := c.FormFile("file")
		checkError(err)
		// Read string from file
		fileString, err := file.Open()
		checkError(err)
		defer fileString.Close()
		dnaSequenceBytes, err := ioutil.ReadAll(fileString)
		checkError(err)
		newDisease.DNASequence = string(dnaSequenceBytes)
		// Capture the disease name and dna sequence.
		diseaseName := newDisease.DiseaseName
		dnaSequence := newDisease.DNASequence
		// Insert the disease into the database.
		_, err = db.Exec("INSERT INTO disease (disease_name, dna_sequence) VALUES (?, ?)", diseaseName, dnaSequence)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Return the disease.
		c.JSON(200, gin.H{"disease_name": diseaseName, "dna_sequence": dnaSequence})
	})

	r.POST("/api/prediction", func(c *gin.Context) {
		// Capture the prediction from the request body.
		var newPrediction Prediction
		if err := c.Bind(&newPrediction); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(500, gin.H{"result": newPrediction})
		// Capture the prediction date, patient name, disease name, result and accuracy.
		patientName := newPrediction.PatientName
		diseasePrediction := newPrediction.DiseasePrediction
		result := newPrediction.Result     // TODO: ganti sesuai algo strmatch
		accuracy := newPrediction.Accuracy // TODO: ganti sesuai algo strmatch
		// Insert the prediction into the database.
		_, err := db.Exec("INSERT INTO prediction (date, patient_name, disease_name, result, accuracy) VALUES (NOW(), ?, ?, ?, ?)", patientName, diseasePrediction, result, accuracy)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Return the prediction.
		c.JSON(200, gin.H{"patient_name": patientName, "disease_name": diseasePrediction, "result": result, "accuracy": accuracy})
	})

	r.DELETE("/api/disease/:id", func(c *gin.Context) {
		// Get the id parameter from the URL.
		id := c.Params.ByName("id")
		// Delete the disease from the database.
		_, err := db.Exec("DELETE FROM disease WHERE id = ?", id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Return the deleted disease.
		c.JSON(200, gin.H{"id": id})
	})

	r.DELETE("/api/prediction/:id", func(c *gin.Context) {
		// Get the id parameter from the URL.
		id := c.Params.ByName("id")
		// Delete the prediction from the database.
		_, err := db.Exec("DELETE FROM prediction WHERE id = ?", id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		// Return the deleted prediction.
		c.JSON(200, gin.H{"id": id})
	})

	r.Run(":8080")
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
