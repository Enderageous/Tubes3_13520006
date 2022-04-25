package main

import (
	"log"
	"io/ioutil"
	"mime/multipart"
)

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func readFile(file *multipart.FileHeader) (string, error) {
	fileString, err := file.Open()
	logError(err)
	defer fileString.Close()
	dnaSequenceBytes, err := ioutil.ReadAll(fileString)
	logError(err)
	return string(dnaSequenceBytes), err
}