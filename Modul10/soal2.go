package main

import "fmt"

func main() {
	var x, y int
	var ikan [1000]float64

	fmt.Scan(&x, &y)

	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var hasil [1000]float64
	jumlahWadah := 0

	i := 0
	for i < x {
		total := 0.0
		count := 0

		for count < y && i < x {
			total += ikan[i]
			i++
			count++
		}

		hasil[jumlahWadah] = total
		jumlahWadah++
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("%.2f ", hasil[i])
	}
	fmt.Println()

	sum := 0.0
	for i := 0; i < jumlahWadah; i++ {
		sum += hasil[i]
	}

	rata := sum / float64(jumlahWadah)

	fmt.Printf("%.2f\n", rata)
}
