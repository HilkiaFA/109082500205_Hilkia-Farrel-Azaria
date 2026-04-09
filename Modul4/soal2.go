package main

import "fmt"

const MAX_SOAL = 8
const TIDAK_SELESAI = 301

func hitungSkor(soal *int, skor *int) {
	var waktu int
	*soal = 0
	*skor = 0

	for i := 0; i < MAX_SOAL; i++ {
		fmt.Scan(&waktu)
		if waktu < TIDAK_SELESAI {
			*soal = *soal + 1
			*skor = *skor + waktu
		}
	}
}

func main() {
	var nama string

	var pemenang string
	var maxSoal, minWaktu int
	minWaktu = 999999

	for {
		fmt.Scan(&nama)

		if nama == "Selesai" {
			break
		}

		var soal, skor int
		hitungSkor(&soal, &skor)

		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}

	fmt.Printf("%s %d %d\n", pemenang, maxSoal, minWaktu)
}
