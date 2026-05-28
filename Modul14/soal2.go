package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func selectionSortAsc(T *arrInt, n int) {
	var i, j, idx_min, temp int

	i = 1
	for i <= n-1 {

		idx_min = i - 1
		j = i

		for j < n {
			if T[idx_min] > T[j] {
				idx_min = j
			}
			j = j + 1
		}

		temp = T[idx_min]
		T[idx_min] = T[i-1]
		T[i-1] = temp

		i = i + 1
	}
}

func selectionSortDesc(T *arrInt, n int) {
	var i, j, idx_max, temp int

	i = 1
	for i <= n-1 {

		idx_max = i - 1
		j = i

		for j < n {
			if T[idx_max] < T[j] {
				idx_max = j
			}
			j = j + 1
		}

		temp = T[idx_max]
		T[idx_max] = T[i-1]
		T[i-1] = temp

		i = i + 1
	}
}

func main() {
	var n, m, x int
	var i, j int

	var ganjil, genap arrInt
	var nGanjil, nGenap int

	fmt.Scan(&n)

	for i = 0; i < n; i++ {

		fmt.Scan(&m)

		nGanjil = 0
		nGenap = 0

		for j = 0; j < m; j++ {

			fmt.Scan(&x)

			if x%2 == 1 {
				ganjil[nGanjil] = x
				nGanjil++
			} else {
				genap[nGenap] = x
				nGenap++
			}
		}

		selectionSortAsc(&ganjil, nGanjil)
		selectionSortDesc(&genap, nGenap)

		for j = 0; j < nGanjil; j++ {
			fmt.Print(ganjil[j], " ")
		}

		for j = 0; j < nGenap; j++ {

			fmt.Print(genap[j])

			if j != nGenap-1 {
				fmt.Print(" ")
			}
		}

		fmt.Println()
		fmt.Println()
	}
}
