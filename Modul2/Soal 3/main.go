package main

import "fmt"

func main() {
	var input, detailkg, detailgr, detailbiayagr, detailbiayakg, total,
		detailbiayagram int
	fmt.Print("Berat Parsel(gram): ")
	fmt.Scan(&input)
	detailkg = input / 1000
	detailgr = input % 1000
	detailbiayakg = detailkg * 10000
	if detailkg > 10 && detailgr < 500 {
		detailbiayagr = 0
		detailbiayagram = detailgr * 15
	} else if detailkg > 10 && detailgr >= 500 {
		detailbiayagr = 0
		detailbiayagram = detailgr * 5
	} else if detailgr < 500 {
		detailbiayagr = detailgr * 15
		detailbiayagram = detailgr * 15
	} else {
		detailbiayagr = detailgr * 5
		detailbiayagram = detailgr * 5
	}
	total = detailbiayakg + detailbiayagr
	fmt.Printf("Detail berat : %d kg + %d gr \n", detailkg, detailgr)
	fmt.Printf("Detail biaya : Rp. %d + Rp. %d\n", detailbiayakg, detailbiayagram)
	fmt.Printf("Total biaya : Rp. %d \n", total)
}
