//Lezione di Mercoledì 14 novembre 2018 

package main 

import (
	."fmt"
)

func main () {
	type anno int // per avere un'alias di Tipo metto un = tra anno e int. la differenza sta che anche se compare come tipo anno avrei comunque una variabile int , con cui potrei fare calcoli con alti tipi int
	var x anno  //non posso assegnare un tipo diverso da anno ad un'altra variabile. se inserissi var y int  e x=y avrei un'errore , è proprio un'altro tipo di variabile come int e float 64
	x = 3 
	Println(x)
}
//l'alias di tipo è comodo per ricordare che tipo abbiamo assegnato ad una variabile più comodamente.

