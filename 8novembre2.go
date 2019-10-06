//Lezione 8 Novembre 2018

package main

import (
	. "fmt"
	"strconv"
)

func rivela(s string) (string, error) {
	t := ""
	for _, r := range s {
		if r != '#' {
			t += string(r)
		}
	}
	x, e := strconv.ParseInt(t, 10, 8)
	if e == nil {
		return strconv.FormatInt(x, 2), nil
	} else {
		return "", e
	}
}

func main() {
	t := ""
	Scan(&t)
	s, e := rivela(t)
	if e == nil {
		println(s)
	} else {
		println(e)
	}
}
