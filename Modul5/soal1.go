package main

import "fmt"

func cepakfibonacci(f int) int {
	if f == 0 {
		return 0
	} else if f == 1 {
		return 1
	}
	return cepakfibonacci(f-1) + cepakfibonacci(f-2)
}

func main() {
	var a int

	fmt.Scan(&a)

	fmt.Println(cepakfibonacci(a))
}
