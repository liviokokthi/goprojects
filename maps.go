/* 22 Novembre 2018
*/

package main 

import ."fmt"
import "os"
/*
func main() {
	m := make(map[int]int) //Dichiarazione mappa , ci metto dentro qualcosa nello stesso modo in cui metto roba dentro le slice
	m[1]=5
	m[0]=2
	//c,ok:= m[1] //ok è un booleano che ti dice se la chiave è stata trovata
	delete(m,0) //Cancella 0,2 
	//Println(c,ok)
	Println(m)
}
*/
//secondo esempio 
/*
func main() {
	m:= map[string]int{"Marco":3 , "Giovanni":2,"Sebastiano":5}
	for nome , punti:= range m{
		Println(nome,punti)
	}
}
*/

//terzo esempio

/*
func main(){
	a:="mamma❤"
	s:= []rune(a)
	t:= string(s)
	Println(a,s,t)
	//NON ANDARE A MODIFICARE I SINGOLI BYTE s[6]='a'
}
*/

//quarto esempio

func main(){
	s:=os.Args[1]
	m:=make(map[rune]string)
	a:=[]rune(s)
	for i := 0 ; i<len(a) -1 ; i++ {
		t,ok := m[a[i]]
		if ok{
			m[a[i]]=t + string(a[i+1])
		} else {
			m[a[i]]=string(a[i+1])
		}
	for k,v := range m {
	Println(string(k),"➡",v)
	}
}
}