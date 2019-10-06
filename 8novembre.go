//Lezione 8 Novembre 2018

package main

import (
	. "fmt"
	"strconv"
)

func main() {
	var s string
	Scan(&s)
	x, e := strconv.ParseInt(s, 2, 0)
	if e == nil { //valore speciale che indica che a "e" non è stato assegnato nessun valore
		t := strconv.Format(int(x), 16)
		Println(t)
	} else {
		println(e)
	}
}
