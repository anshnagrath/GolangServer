package main
import (
	 "fmt"
)
func main( ){
	s:= []string{ "one" ,"two","one","three","five","three","seven","two"}
	ms:= make(map[string]bool)
	for _,v:=range s{
		if !ms[v] {
			ms[v] = true ;
		}
		    
	}
	fmt.Println(ms)
	ns:= make([]string,len(ms))
	for key,val := range ms{
		ns= append(ns,key)
		fmt.Println(val)
	}
	fmt.Println(ns)

}