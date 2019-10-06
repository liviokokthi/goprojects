/* Cicli annidati MOD5 - ALBAN: 

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

    var i, j, k, n int

    Print("Inserire il numero di asterischi:_")
    Scan(&n)

    for i=0; i<n; i++ {

         for j=0; j<n-(i+1); j++{

          Print (" ")

         }

         for k=0; k<2*i+1; k++ {

          Print ("*")

         }

         Println()
    }

      cont:=1

      contk:=0

    for i=n-1; i>0; i-- {
          
         for j=0; j<cont && cont < n; j++{

          Print (" ")

         }

         for k=n-contk; k>0; k-- {

          Print ("*")

         }

       cont++
       contk+=2


         Println()
    }
}