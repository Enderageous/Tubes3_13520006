package main

import (
	"log"
	"io/ioutil"
	"mime/multipart"
	"regexp"
	"strings"
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

func findMonthNumberFromName(monthName string) string {
	months := map[string]string{
		"januari": "01",
		"februari": "02",
		"maret": "03",
		"april": "04",
		"mei": "05",
		"juni": "06",
		"juli": "07",
		"agustus": "08",
		"september": "09",
		"oktober": "10",
		"november": "11",
		"desember": "12",
		"jan": "01",
		"feb": "02",
		"mar": "03",
		"apr": "04",
		"jun": "06",
		"jul": "07",
		"aug": "08",
		"agt": "08",
		"sep": "09",
		"sept": "09",
		"okt": "10",
		"nov": "11",
		"des": "12",
	}
	return months[monthName]
}

func parseDate(s string) (string, string) {
	var fs string
	var format = regexp.MustCompile(`\d{4}\W\d{2}\W\d{2}|\d{2}\W\d{2}\W\d{4}|\d{2}\s\w+\s\d{4}`)
	var format1 = regexp.MustCompile(`\d{4}\W\d{2}\W\d{2}`) // YYYY-MM-DD
	var format2 = regexp.MustCompile(`\d{2}\W\d{2}\W\d{4}`) // DD-MM-YYYY
	var format3 = regexp.MustCompile(`\d{2}\s\w+\s\d{4}`) // DD MONTH YYYY
	result := format.FindString(s)
	if result != "" {
		fs = format.ReplaceAllString(s, "")
		if format1.FindString(result) != "" {
			result = format1.FindString(result)
			// Change result format to YYYY-MM-DD
			result = result[0:4] + "-" + result[4:6] + "-" + result[6:8]
			return result, fs
		} else if format2.FindString(result) != "" {
			result = format2.FindString(result)
			// Change result format from DD-MM-YYYY to YYYY-MM-DD
			result = result[6:10] + "-" + result[3:5] + "-" + result[0:2]
			return result, fs
		} else if format3.FindString(result) != "" {
			result = format3.FindString(result)
			// Find month name
			monthName := regexp.MustCompile(`[a-zA-Z]+`).FindString(result)
			// Change result format from DD MONTH YYYY to YYYY-MM-DD
			result = result[len(result)-4:] + "-" + findMonthNumberFromName(strings.ToLower(monthName)) + "-" + result[0:2]
			return result, fs
		}
		return result, fs
	}
	return "", s
}

func parseWord(s string) string{
	var result = regexp.MustCompile(`[a-zA-Z]+`).FindString
	return result(s)
}