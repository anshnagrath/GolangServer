package main

import (
	"fmt"
	"sync"
)

func main() {
	//lamba/anyomous  function for quick use and discard
	// wait groups
	wg := &sync.WaitGroup{}
	//Adding number of go routines to a wait group
	wg.Add(2)
	go func() {
		fmt.Println("Lamda function envoked inside  go routine 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("Lamda2 function envoked inside go routine 2 ")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("outside routinev")

}
