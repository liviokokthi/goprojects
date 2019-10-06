/* Cicli annidati MOD3: 

*****
 ****
  ***                   
   **
    *              
                   */

  package main 

  import (
            ."fmt"
                    )

  func main(){

  	var i, j, k, n int

    Print("Inserire il numero di asterischi:_")
    Scan(&n)

    for i=0; i<n; i++{

        for j=0; j<i; j++{

            Print(" ")
        }

        for k=0; k<n-i; k++{

          Print("*")
        }

        Println()
    }
}
