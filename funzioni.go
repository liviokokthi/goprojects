//7 novembre 2018
//funzioni esterne al main
//ripreso codice Collatz

package main 

import(
."fmt"
) 

//PRIMO ESERCIZIO


func collatz (x int) int{
if x%2 ==o {
return x/2 
} else {
return 3*x+1
}
}


func main() {
var n int 
Scan(&n)

for n !=1 {
Println(n)
n=collatz(n) //chiamata funzione , tra parentesi c'� arogmento

}

}



//ESERCIZIO +1 

func oneplus(x int) {
x++
return
}

func main() {
var n int 
Scan(&n)
oneplus(n+9)
Println(n)
}


//ESERCIZIO STAMPA CIAO

func stampaciao(x int) {
for i:=0;i<x,i++{
Println("Ciao")
}
return
}

func main () {
var n int 
for {
Scan(&n)
stampaciao(n*n)
}
}


//ESERCIZIO SOMMA1

func stampaciao(x int) {
for i:=0;i<x,i++{
Println("Ciao")
}
return
}

func main () {
var n int 
for {
Scan(&n)
stampaciao(n*n)
}
}























