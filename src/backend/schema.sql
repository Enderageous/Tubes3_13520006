CREATE TABLE disease (
	disease_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	disease_name VARCHAR(50) NOT NULL,
	dna_sequence VARCHAR(255) NOT NULL
);

CREATE TABLE prediction (
	prediction_id INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
	prediction_date DATE NOT NULL,
	patient_name VARCHAR(50) NOT NULL,
	dna_sequence VARCHAR(255) NOT NULL,
	disease_id INT NOT NULL,
	result BOOLEAN NOT NULL,
	accuracy FLOAT NOT NULL,
	FOREIGN KEY (disease_id) REFERENCES disease(disease_id),
	CONSTRAINT accuracy CHECK (accuracy >= 0 and accuracy <= 1)
);

CREATE VIEW prediction_view AS
	SELECT prediction_id, prediction_date, patient_name, 
	prediction.dna_sequence, prediction.disease_id, disease_name, result, accuracy
	FROM prediction, disease 
	WHERE prediction.disease_id=disease.disease_id;