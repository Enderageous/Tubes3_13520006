package main

type Disease struct {
    ID          int    `json:"id" form:"id"`
    Name        string `json:"disease_name" form:"disease_name"`
    DNASequence string `json:"dna_sequence" form:"dna_sequence"`
}

type Prediction struct {
    ID          int     `json:"id" form:"id"`
    Date        string  `json:"date" form:"date"`
    PatientName string  `json:"patient_name" form:"patient_name"`
    DNASequence string  `json:"dna_sequence" form:"dna_sequence"`
    DiseaseId   int     `json:"disease_id" form:"disease_id"`     // foreign key ke disease(id)
    DiseaseName string `json:"disease_name" form:"disease_name"`  // ambil dari prediction_view
    Result      bool    `json:"result" form:"result"`             // 0 = false, 1 = true
    Accuracy    float32 `json:"accuracy" form:"accuracy"`         // 0 <= accuracy <= 1
}
