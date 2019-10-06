//Lezione di Mercoledì 14 novembre 2018 

package main 

import ( 
	."fmt"
)

type data struct {
	g,m,a int 
}
type nascitamorte struct {
	nascita, morte data
}

type persona struct {
	nome string 
	nm nascitamorte 
}

/*
func natale (anno int ) data{
	p:=new(data)
	*p data {21,12,anno}
	// posso ottenere tutta sta roba ugualmente scrivendo qua dentro solo: return data {25,12,anno}
	return *p
}
*/

func main () {
	var p persona 

	//p.nm.nascita.g = data { 25, 12, 1967}

	//nel caso in cui volessi assegnare entrambe le date 
	p.nm = nascitamorte {data{1,1,1980},data{1,1,2050}}
	Println(p)
}
