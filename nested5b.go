/* Cicli annidati MOD5 - PROF. : 

  *               
 ***
*****
 ***
  *                   
                  */

  package main 

  import (
            ."fmt"
                    )

  func main(){

    var n int

    for{

    Print("Inserire il numero di asterischi:_")
    Scan (&n)	

    if n%2==0{

    	Println ("Errore(!)",n,"è un numero pari")
    	continue 
    } else { 
    	break }
 }

    for i:=0; i<=n/2; i++ {

    	for j:=0; j<n/2-i; j++{

    		Print (" ")
    	}

    	for j:=0; j<2*i+1; j++ {

    		Print ("*")
    	}

    	Println()

    }
    
    for i:=n/2-1; i>=0; i--{

    	for j:=0; j<n/2-i; j++{

    		Print(" ")
    	}

    	for j:=0; j<2*i+1; j++{

    		Print ("*")
    	}

    	Println()
      }
    }
  
