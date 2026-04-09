package main

import "fmt"

func bagi(x, simpan int) {

	if x == simpan {
		fmt.Print(x)
	} else {
		if simpan%x == 0 {
			fmt.Print(x, " ")
		}

		bagi(x+1, simpan)
	}
}

func main() {
	var n int
	fmt.Scan(&n)

	bagi(1, n)
}
