package main

import (
		 ."fmt"
		 "os"
		 "bufio"
		 "strings"
				)
			

func main(){
	scanner:= bufio.NewScanner(os.Stdin)
	doc := []string{}
	for scanner.Scan(){
		doc=append(doc,scanner.Text())
	}
	mr := make(map[string][]int)

	for i,d := range doc{
		d = strings.Replace(d, ".", " ",-1)
		d = strings.Replace(d, ",", " ",-1)
		a := strings.Split(d," ")

		for _,p := range a {
			if len(p) ==0 {continue}
			mr[p] = append(mr[p],i)
		}
	}
	
	q:=os.Args[1]
	r:=os.Args[2]

	mq:=map[int]bool{}
	for _,d:= range mr[q] {
		mq[d]=true
	}

	for _, d:= range mr[r] {
		if ! mq[d] { continue }
		s:= doc[d]
		j:=strings.Index(s,q)
		Print(d,"▪", s[:j],"\033[31;1;4m",q,"\033[0m",s[j + len(q)],'\n')
	}

	if i < len(s) { return s[i] } 
	
	
}