package main

type Disease struct {
	ID          int    `json:"id" form:"id"`
	DiseaseName string `json:"disease_name" form:"disease_name"`
	DNASequence string `json:"dna_sequence" form:"dna_sequence"`
}

type Prediction struct {
	ID                int     `json:"id" form:"id"`
	Date              string  `json:"date" form:"date"`
	PatientName       string  `json:"patient_name" form:"patient_name"`
	DiseasePrediction string  `json:"disease_prediction" form:"disease_prediction"`
	Result            bool    `json:"result" form:"result"`
	Accuracy          float64 `json:"accuracy" form:"accuracy"`
}
