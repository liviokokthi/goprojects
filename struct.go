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
func invecchia (p *persona) {

	p.eta++ //funziona anche con solo la p , semplificazione di go  anche se non consigliato perchè ci si potrebbe confondere 

}


func main () {

	var x,y persona


	x.nome = "Livio"
	x.cognome ="Kokthi"
	x.eta = 22

	y.nome = "C"
	y.cognome ="R"
	y.eta = 7
	
	//per velocizzare facciamo:    x := persona {nome:"Livio" ,cognome:"Kokthi" ,eta: 22} 
	//con questo metodo sopra posso avere il := dove la variabile x riconosce automaticamente il tipo persona e lo assegna a x
	//posso mettere i campi che voglio , se non li metto verrà stampat auna stringa vuota. 
	//l'ordine con cui metto i campi non è fondamentale  e posso inserirli quando e come voglio.
	//ESEMPIO  x := persona {cognome: "kokthi" , eta: 22 }
	invecchia(&x)

	Println (x,"x",y)
}


/*
uso p diretammente  

{
	p:= new(persona)
	*p = persona {cognome:"kokthi" , nome: "livio",era: 22}
	invecchia(p)
	Println(p,*p)

}