package main

import "fmt"

func Star(total int) {
	if total == 0 {
		return
	}
	fmt.Print("*")
	Star(total - 1)
}

func Pattern(n int, b int) {
	if b > n {
		return
	}

	Star(b)
	fmt.Println()

	Pattern(n, b+1)
}

func main() {
	var n int
	fmt.Scan(&n)

	Pattern(n, 1)
}
