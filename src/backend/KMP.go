package main

import (
	"fmt"
	"regexp"
)

// Check apakah input file hanya memiliki A, G, C, T, dan tidak ada spasi
func isAGCT(s string) bool{
	var result = regexp.MustCompile(`^[ACGT]+$`).MatchString

	return (result(s))
}

// Algoritma Utama KMP
func KMP(input string, disease string) bool {
	border := computeBorder(disease)
	i := 0
	j := 0
	m := len(disease)
	n := len(input)
	found := false

	for i < n && !found{
		fmt.Println(i, j)
		if input[i] == disease[j] {
			if j == m-1 {
				fmt.Println(i - m + 1)
				found = true
			}
			i++
			j++
		} else if (j > 0) {
			j = border[j-1]
		} else {
			i++
		}
	}
	return found
}

// Fungsi untuk menghitung border function
func computeBorder(input string) []int {
	var border = make([]int, len(input))
	i := 1
	j := 0
	border[0] = 0

	for i < len(input) {
		if input[i] == input[j] {
			border[i] = j + 1
			i++
			j++
		} else if (j > 0) {
			j = border[j-1]
		} else {
			border[i] = 0
			i++
		}
	}
	fmt.Println(border)
	return border
}

// Fungsi untuk menggunakan Algoritma KMP
func mainKMP(Input string, Disease string) bool {
	var result bool

	if isAGCT(Input) && isAGCT(Disease) {
		result = KMP(Input, Disease)
	} else {
		fmt.Println("Input format is not A, C, G, T, or it has spaces")
	}

	return result
}