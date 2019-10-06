//Lezione 8 Novembre 2018

package main

import (
	. "fmt"
)

func primo(n int) bool {
	for d := 2; d < n; d++ {
		if n%d == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	Scan(&n)
	prec := false
	for i := 3; i < n; i += 2 {
		t := primo(i)
		if t && prec {
			Println(i-2, i)
		}
		prec = t
	}

}
