// Strings

/* rune : un elemento int32bit, equivalente del "char*"

  x = 'x'
  x = '\n'

Println (x) : prende il valore HEX di x

x = '\u2264'

Println(x,string(x)) */

// Stampa una stringa

package main

import (
	. "fmt"
)

func main() {

	var s string
	var x rune

	Print("Inserire una Stringa: _ ")
	Scan(&s)

	x = '\u2264'

	Println(s, x)
	Println(len(s))
}
