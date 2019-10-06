/*
Tipi base  : complex, int, string ...
Tipi Composti : Strutture 
Tipi Interfaccia 

 i tipi puntatore rappresentantano un'astrazione di numeri che indicano dove si trovano in memoria i dati.
-tramite i puntatori è possibile fare un riferimento su milioni di interi evitando di fare una copia.
in go l'uso dei puntatori vengono descitti mettendo "*"" : esempio var p *int/string/uint8...

nil rappresenta un valore di puntatore che non punta a nulla. (puntatore non valido) e posso assegnarlo ad ogni variabile puntatore
esempio p=nil 
l'operatore & da una posizione in memoria della variabile.

var p *int
var c in
p=&x

l'operatore *  è l'operatore di indirezione 

*p=0  ; metto 0 dentro x , non dentro p perchè è la variabile che ho puntato

p=nil NON POSSO FARLO , VIENE ERRORE. 

quindi ricapitolando i 3 simboli sono :   * , & ,nil

*/

//PRIMO ESEMPIO

 
package main 

import (
            ."fmt" 
                    )

func main (){
var s string = "ciao"
var p *string = &s
t:=""
Println(p,s)
p = &t //leggi sotto , con &t ritorno a alla stringa "ciao"
*p="mamma"
Println(s) // *p assegnato prima ora è diventato s
Println((*p)[0]) // [0]indica la posizione in memoria
// *p diventa s e stamperò ciao
}
*/

/*
 x:=5
p:=&x(*int)
var q **int=&p
**q=3

*/

//SECONDO ESEMPIO

package main 

import (
            ."fmt" 
                    )

func più1(p*int) {
	(*p)++
}

func main () {
	var p *int
	p= new(int)
	*p = 5
	q := p 
	Println(*p,p)
	più1(q)
	Println(*p,p)
}

