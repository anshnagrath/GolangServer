package main
import (
	"fmt"
	"time"
)

func main(){
	//unbuffered chan can hold only one value
	c:= make(chan int)
	fmt.Println("Channels",c)
	go func(){
		c <-1
	}()
	val := <- c
	fmt.Println(val,"value from the channel")	
	go func(){
		c <-2
	}()

	time.Sleep(time.Second * 5)
	//val = <-c
	fmt.Println(val,c,"value from channel")
}