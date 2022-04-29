package main

// Algoritma Utama Boyer-Moore
func BMS(input []byte, disease []byte) bool {
	last := lastOcc(disease)
	m := len(disease)
	n := len(input)
	i := m-1
	j := m -1
	found := false

	if i <= n-1 {
		for i <= n-1 && !found {
			if input[i] == disease[j] {
				if j == 0 {
					found = true
				} else {
					i--
					j--
				}
			} else {
				jump := last[input[i]]
				if (j < 1+jump) {
					i = i + m - j
				} else {
					i = i + m - (1 + jump)
				}
				j = m - 1
			}
		}
	}

	return found
}

// Fungsi untuk membentuk sebuah array yang berisi kemunculan terakhir setiap karakter ASCII dalam penyakit
func lastOcc(disease []byte) [91]int {
	var last [91]int
	i := 0

	for i < 91 {
		last[i] = -1
		i++
	}

	i = 0
	for i < len(disease) {
		last[disease[i]] = i
		i++
	}

	return last
}

// Fungsi untuk menggunakan Algoritma Boyer-Moore
func mainBMS(Input string, Disease string) bool {
	var result bool

	InputByte := []byte(Input)
	DiseaseByte := []byte(Disease)

	result = BMS(InputByte, DiseaseByte)
	return result
}