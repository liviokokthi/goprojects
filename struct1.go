//Lezione di Mercoledì 14 novembre 2018 

package main 

import ( 
	."fmt"
)

type persona struct {
	nome string
	cognome string
	eta int
}
type matricola struct {
	persona persona
	id int 
}

func invecchia (q*persona) {
	q.eta++
}

func main () {
	var x persona
	Scan(&x.nome)
	Scan(&x.cognome)
	Scan(&x.eta)
	Println(x)
}