package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0

	fmt.Print("Teks : ")

	for {
		fmt.Scanf("%c", &ch)

		if ch == '\n' {
			break
		}

		if ch == '\r' {
			continue
		}

		if ch == ' ' {
			continue
		}

		t[*n] = ch
		*n++
	}
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	var temp tabel

	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}

	balikanArray(&temp, n)

	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Palindrom ? ")
	fmt.Println(palindrom(tab, n))
}
