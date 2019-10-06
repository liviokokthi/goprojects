 // range of strings

 package main

 import (
           ."fmt"
                   )

 func main(){

 	s:=" "

 	Print("Inserire una stringa: _ ")
 	Scan(&s)

 	for i, r := range s{

 		Println (i,r, string(r))
 	}
 }