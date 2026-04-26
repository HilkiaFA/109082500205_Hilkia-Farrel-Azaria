package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int

	var pemenang [100]string
	var n int = 0

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	j := 1

	for {
		fmt.Printf("Pertandingan %d : ", j)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[n] = klubA
			n++
		} else if skorB > skorA {
			pemenang[n] = klubB
			n++
		} else {
			pemenang[n] = "Draw"
			n++
		}

		j++
	}

	for i := 0; i < j-1; i++ {
		fmt.Printf("Hasil %d : %v\n", i+1, pemenang[i])

	}
	fmt.Println("Pertandingan selesai")
}
