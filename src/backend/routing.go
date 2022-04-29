package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

func TestAPI(c *gin.Context) {
	c.String(200, "200 OK")
}

func GetDisease(c *gin.Context) {
	// Get all diseases from the database.
	rows, err := db.Query("SELECT disease_id, disease_name, dna_sequence FROM disease")
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
	err := db.QueryRow("SELECT disease_id, disease_name, dna_sequence FROM disease WHERE disease_id = ?", id).Scan(&disease.ID, &disease.Name, &disease.DNASequence)
	getError(c, err)
	// Return the disease.
	c.JSON(200, disease)
}

func GetPrediction(c *gin.Context) {
	q := c.Query("q")
	if q != "" {
		// Parse params
		date, q := parseDate(q)
		word := parseWord(q)
		var rows *sql.Rows
		var err error
		if date != "" && word != "" {
			rows, err = db.Query("SELECT prediction_id, prediction_date, patient_name, prediction.dna_sequence, prediction.disease_id, disease_name, result, accuracy FROM prediction, disease WHERE prediction.disease_id=disease.disease_id and prediction_date = ? and disease_name LIKE ?", date, "%"+word+"%")
		} else if date != "" {
			rows, err = db.Query("SELECT prediction_id, prediction_date, patient_name, prediction.dna_sequence, prediction.disease_id, disease_name, result, accuracy FROM prediction, disease WHERE prediction.disease_id=disease.disease_id AND prediction_date = ?", date)
		} else if word != "" {
			rows, err = db.Query("SELECT prediction_id, prediction_date, patient_name, prediction.dna_sequence, prediction.disease_id, disease_name, result, accuracy FROM prediction, disease WHERE prediction.disease_id=disease.disease_id AND disease_name LIKE ?", "%"+word+"%")
		}
		getError(c, err)
		defer rows.Close()
		// Iterate through the predictions and create a JSON array.
		var predictions []Prediction
		for rows.Next() {
			var prediction Prediction
			err := rows.Scan(&prediction.ID, &prediction.Date, &prediction.PatientName, &prediction.DNASequence, &prediction.DiseaseId, &prediction.DiseaseName, &prediction.Result, &prediction.Accuracy)
			getError(c, err)
			predictions = append(predictions, prediction)
		}
		// Return the predictions.
		c.JSON(200, predictions)
	} else {
		// Get all predictions from the database.
		rows, err := db.Query("SELECT prediction_id, prediction_date, patient_name, prediction.dna_sequence, prediction.disease_id, disease_name, result, accuracy FROM prediction, disease WHERE prediction.disease_id=disease.disease_id")
		getError(c, err)
		defer rows.Close()
		// Iterate through the predictions and create a JSON array.
		var predictions []Prediction
		for rows.Next() {
			var prediction Prediction
			err := rows.Scan(&prediction.ID, &prediction.Date, &prediction.PatientName, &prediction.DNASequence, &prediction.DiseaseId, &prediction.DiseaseName, &prediction.Result, &prediction.Accuracy)
			getError(c, err)
			predictions = append(predictions, prediction)
		}
		// Return the predictions.
		c.JSON(200, predictions)
	}
}

func GetPredictionById(c *gin.Context) {
	// Get the id parameter from the URL.
	id := c.Params.ByName("id")
	// Get the prediction from the database.
	var prediction Prediction
	err := db.QueryRow("SELECT prediction_id, prediction_date, patient_name, prediction.dna_sequence, prediction.disease_id, disease_name, result, accuracy FROM prediction, disease WHERE prediction.disease_id=disease.disease_id and prediction.prediction_id = ?", id).Scan(&prediction.ID, &prediction.Date, &prediction.PatientName, &prediction.DNASequence, &prediction.DiseaseId, &prediction.DiseaseName, &prediction.Result, &prediction.Accuracy)
	getError(c, err)
	// Return the prediction.
	c.JSON(200, prediction)
}

func PostDisease(c *gin.Context) {
	// Capture the disease from the request body.
	var newDisease Disease
	err := c.Bind(&newDisease)
	getError(c, err)
	file, err := c.FormFile("dna_file")
	logError(err)

	newDisease.DNASequence, err = readFile(file)
	logError(err)

	if (isAGCT(newDisease.DNASequence)) {
		// Capture the disease name and dna sequence.
		diseaseName := newDisease.Name
		dnaSequence := newDisease.DNASequence
		// Insert the disease into the database.
		_, err = db.Exec("INSERT INTO disease (disease_name, dna_sequence) VALUES (?, ?)", diseaseName, dnaSequence)
		getError(c, err)
		// Return the disease.
		c.JSON(200, gin.H{"disease_name": diseaseName, "dna_sequence": dnaSequence})
	} else {
		c.JSON(400, gin.H{"error": "Invalid DNA sequence"})
	}
}

func PostPrediction(c *gin.Context) {
	// Capture the prediction from the request body.
	var newPrediction Prediction
	err := c.Bind(&newPrediction)
	getError(c, err)
	file, err := c.FormFile("dna_file")
	logError(err)
	
	newPrediction.DNASequence, err = readFile(file)
	logError(err)

	if (isAGCT(newPrediction.DNASequence)) {
		// Capture the prediction date, patient name and disease name
		patientName := newPrediction.PatientName
		diseaseId := newPrediction.DiseaseId

		// Find the disease DNA sequence from the database.
		diseaseDNARow, err := db.Query("SELECT dna_sequence FROM disease WHERE disease_id = ?", diseaseId)
		getError(c, err)
		defer diseaseDNARow.Close()
		var diseaseDNASequence string
		for diseaseDNARow.Next() {
			err := diseaseDNARow.Scan(&diseaseDNASequence)
			getError(c, err)
		}

		dnaSequence := newPrediction.DNASequence
		newPrediction.Result = mainKMP(dnaSequence, diseaseDNASequence)
		result := newPrediction.Result
		accuracy := 1 // TODO: ganti sesuai algo strmatch
		// Insert the prediction into the database.
		_, err = db.Exec("INSERT INTO prediction (prediction_date, patient_name, dna_sequence, disease_id, result, accuracy) VALUES (NOW(), ?, ?, ?, ?, ?)", patientName, dnaSequence, diseaseId, result, accuracy)
		getError(c, err)

		// get ID of recently inserted prediction
		err = db.QueryRow("SELECT prediction_id, prediction_date, disease_name FROM prediction, disease WHERE prediction.disease_id=disease.disease_id AND patient_name = ? AND prediction.dna_sequence = ? AND prediction.disease_id = ?", patientName, dnaSequence, diseaseId).Scan(&newPrediction.ID, &newPrediction.Date, &newPrediction.DiseaseName)
		getError(c, err)

		// Return the prediction.
		c.JSON(200, newPrediction)
	} else {
		c.JSON(400, gin.H{"error": "Invalid DNA sequence"})
	}
}

func DeleteDiseaseById(c *gin.Context) {
	// Get the id parameter from the URL.
	id := c.Params.ByName("id")
	// Delete the disease from the database.
	_, err := db.Exec("DELETE FROM disease WHERE disease_id = ?", id)
	getError(c, err)
	// Return the deleted disease.
	c.JSON(200, gin.H{"id": id})
}

func DeletePredictionById(c *gin.Context) {
	// Get the id parameter from the URL.
	id := c.Params.ByName("id")
	// Delete the prediction from the database.
	_, err := db.Exec("DELETE FROM prediction WHERE prediction_id = ?", id)
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