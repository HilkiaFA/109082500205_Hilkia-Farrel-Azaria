package main

import "fmt"

const NMAX int = 1000000

type arrInt [NMAX]int

func selectionSort(T *arrInt, n int) {

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

func main() {
	var data arrInt
	var x, n int
	var median int

	n = 0

	for {

		fmt.Scan(&x)

		if x == -5313 {
			break
		}

		if x != 0 {

			data[n] = x
			n++

		} else {

			selectionSort(&data, n)

			if n%2 == 1 {
				median = data[n/2]
			} else {
				median = (data[(n/2)-1] + data[n/2]) / 2
			}

			fmt.Println(median)
		}
	}
}
