package main

import (
	"regexp"
)

// Check apakah input file hanya memiliki A, G, C, T, dan tidak ada spasi
func isFileAGCT(s []byte) bool{
	var result = regexp.MustCompile(`^[ACGT]+$`).MatchString

	return (result(string(s)))
}