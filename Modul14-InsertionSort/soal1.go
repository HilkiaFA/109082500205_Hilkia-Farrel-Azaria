package main

import "fmt"

const NMAX int = 1000

type arrInt [NMAX]int

func bacaData(A *arrInt, n *int) {
	var x int

	*n = 0

	fmt.Scan(&x)
	for x >= 0 {
		A[*n] = x
		*n++

		fmt.Scan(&x)
	}
}

func insertionSort(A *arrInt, n int) {
	var i, j, temp int

	i = 1
	for i <= n-1 {
		j = i
		temp = A[j]

		// ascending
		for j > 0 && temp < A[j-1] {
			A[j] = A[j-1]
			j--
		}

		A[j] = temp
		i++
	}
}

func cekJarak(A arrInt, n int) {
	var jarak int
	var tetap bool
	var i int

	if n <= 1 {
		fmt.Println("Data berjarak 0")
		return
	}

	jarak = A[1] - A[0]
	tetap = true

	i = 2
	for i < n && tetap {
		if A[i]-A[i-1] != jarak {
			tetap = false
		}
		i++
	}

	if tetap {
		fmt.Println("Data berjarak", jarak)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

func cetakArray(A arrInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(A[i])

		if i < n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	var A arrInt
	var n int

	bacaData(&A, &n)

	insertionSort(&A, n)

	cetakArray(A, n)

	cekJarak(A, n)
}
