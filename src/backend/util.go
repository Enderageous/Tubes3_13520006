package main

import (
	"log"
	"io/ioutil"
	"mime/multipart"
	"regexp"
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

func parseDate(s string) string{
	var result = regexp.MustCompile(`\d{4}-\d{2}-\d{2}`).FindString
	return result(s)
}

func parseWord(s string) string{
	var result = regexp.MustCompile(`[a-zA-Z]+`).FindString
	return result(s)
}