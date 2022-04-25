package main

import (
	"io/ioutil"
	"github.com/gin-gonic/gin"
)

func TestAPI(c *gin.Context) {
	c.String(200, "200 OK")
}

func GetDisease(c *gin.Context) {
	// Get all diseases from the database.
	rows, err := db.Query("SELECT id, name, dna_sequence FROM disease")
	getError(c, err)
	defer rows.Close()
	// Iterate through the diseases and create a JSON array.
	var diseases []Disease
	for rows.Next() {
		var disease Disease
		err := rows.Scan(&disease.ID, &disease.Name, &disease.DNASequence)
		getError(c, err)
		diseases = append(diseases, disease)
	}
	// Return the diseases.
	c.JSON(200, diseases)
}

func GetDiseaseById(c *gin.Context) {
	// Get the id parameter from the URL.
	id := c.Params.ByName("id")
	// Get the disease from the database.
	var disease Disease
	err := db.QueryRow("SELECT id, name, dna_sequence FROM disease WHERE id = ?", id).Scan(&disease.ID, &disease.Name, &disease.DNASequence)
	getError(c, err)
	// Return the disease.
	c.JSON(200, disease)
}

func GetPrediction(c *gin.Context) {
	// Get all predictions from the database.
	rows, err := db.Query("SELECT id, date, patient_name, disease_id, disease_name, result, accuracy FROM prediction_view")
	getError(c, err)
	defer rows.Close()
	// Iterate through the predictions and create a JSON array.
	var predictions []Prediction
	for rows.Next() {
		var prediction Prediction
		err := rows.Scan(&prediction.ID, &prediction.Date, &prediction.PatientName, &prediction.DiseaseId, &prediction.DiseaseName, &prediction.Result, &prediction.Accuracy)
		getError(c, err)
		predictions = append(predictions, prediction)
	}
	// Return the predictions.
	c.JSON(200, predictions)
}

func GetPredictionById(c *gin.Context) {
	// Get the id parameter from the URL.
	id := c.Params.ByName("id")
	// Get the prediction from the database.
	var prediction Prediction
	err := db.QueryRow("SELECT id, date, patient_name, disease_id, disease_name, result, accuracy FROM prediction_view WHERE id = ?", id).Scan(&prediction.ID, &prediction.Date, &prediction.PatientName, &prediction.DiseaseId, &prediction.DiseaseName, &prediction.Result, &prediction.Accuracy)
	getError(c, err)
	// Return the prediction.
	c.JSON(200, prediction)
}

func InsertDisease(c *gin.Context) {
	// Capture the disease from the request body.
	var newDisease Disease
	err := c.Bind(&newDisease)
	getError(c, err)
	file, err := c.FormFile("file")
	logError(err)
	// Read string from file
	fileString, err := file.Open()
	logError(err)
	defer fileString.Close()
	dnaSequenceBytes, err := ioutil.ReadAll(fileString)
	logError(err)
	newDisease.DNASequence = string(dnaSequenceBytes)
	// Capture the disease name and dna sequence.
	diseaseName := newDisease.Name
	dnaSequence := newDisease.DNASequence
	// Insert the disease into the database.
	_, err = db.Exec("INSERT INTO disease (name, dna_sequence) VALUES (?, ?)", diseaseName, dnaSequence)
	getError(c, err)
	// Return the disease.
	c.JSON(200, gin.H{"disease_name": diseaseName, "dna_sequence": dnaSequence})
}

func InsertPrediction(c *gin.Context) {
	// Capture the prediction from the request body.
	var newPrediction Prediction
	err := c.Bind(&newPrediction)
	getError(c, err)

	file, err := c.FormFile("file")
	logError(err)
	// Read string from file
	fileString, err := file.Open()
	logError(err)
	defer fileString.Close()
	dnaSequenceBytes, err := ioutil.ReadAll(fileString)
	logError(err)
	newPrediction.DNASequence = string(dnaSequenceBytes)

	// Capture the prediction date, patient name, disease name, result and accuracy.
	patientName := newPrediction.PatientName
	diseaseId := newPrediction.DiseaseId
	result := 1   // TODO: ganti sesuai algo strmatch
	accuracy := 1 // TODO: ganti sesuai algo strmatch
	// Insert the prediction into the database.
	_, err = db.Exec("INSERT INTO prediction (date, patient_name, disease_id, result, accuracy) VALUES (NOW(), ?, ?, ?, ?)", patientName, diseaseId, result, accuracy)
	getError(c, err)
	// Return the prediction.
	c.JSON(200, gin.H{"patient_name": patientName, "disease_id": diseaseId, "result": result, "accuracy": accuracy})
}

func DeleteDiseaseById(c *gin.Context) {
	// Get the id parameter from the URL.
	id := c.Params.ByName("id")
	// Delete the disease from the database.
	_, err := db.Exec("DELETE FROM disease WHERE id = ?", id)
	getError(c, err)
	// Return the deleted disease.
	c.JSON(200, gin.H{"id": id})
}

func DeletePredictionById(c *gin.Context) {
	// Get the id parameter from the URL.
	id := c.Params.ByName("id")
	// Delete the prediction from the database.
	_, err := db.Exec("DELETE FROM prediction WHERE id = ?", id)
	getError(c, err)
	// Return the deleted prediction.
	c.JSON(200, gin.H{"id": id})
}

func getError(c *gin.Context, err error) {
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
}