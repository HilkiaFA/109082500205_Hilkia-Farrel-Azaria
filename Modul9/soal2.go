package main

import "fmt"

func sqrt(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := x
	for i := 0; i < 20; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen (N): ")
	fmt.Scan(&n)

	var arr [100]int

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSemua elemen:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	fmt.Println("\nIndeks ganjil:")
	for i := 0; i < n; i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	fmt.Println("\nIndeks genap:")
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var x int
	fmt.Print("\nMasukkan x: ")
	fmt.Scan(&x)

	fmt.Println("Indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	var idx int
	fmt.Print("\nMasukkan indeks yang dihapus: ")
	fmt.Scan(&idx)

	for i := idx; i < n-1; i++ {
		arr[i] = arr[i+1]
	}
	n--

	fmt.Println("Array setelah dihapus:")
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	var sum int
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	avg := float64(sum) / float64(n)
	fmt.Println("\nRata-rata:", avg)

	var variance float64
	for i := 0; i < n; i++ {
		diff := float64(arr[i]) - avg
		variance += diff * diff
	}
	variance = variance / float64(n)

	stdDev := sqrt(variance)
	fmt.Println("Standar deviasi:", stdDev)

	var cari int
	fmt.Print("\nMasukkan nilai yang dicari: ")
	fmt.Scan(&cari)

	freq := 0
	for i := 0; i < n; i++ {
		if arr[i] == cari {
			freq++
		}
	}
	fmt.Println("Frekuensi:", freq)
}
