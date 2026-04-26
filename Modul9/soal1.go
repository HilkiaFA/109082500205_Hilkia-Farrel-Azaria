package main

import "fmt"

type titik struct {
	x int
	y int
}

type lingkaran struct {
	pusat titik
	r     int
}

func jarak(p, q titik) int {
	dx := p.x - q.x
	dy := p.y - q.y
	return dx*dx + dy*dy
}

func didalam(c lingkaran, p titik) bool {
	return jarak(p, c.pusat) < c.r*c.r
}

func main() {
	var c1, c2 lingkaran
	var p titik

	fmt.Scan(&c1.pusat.x, &c1.pusat.y, &c1.r)

	fmt.Scan(&c2.pusat.x, &c2.pusat.y, &c2.r)

	fmt.Scan(&p.x, &p.y)

	in1 := didalam(c1, p)
	in2 := didalam(c2, p)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
