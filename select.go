package main

import (
	"fmt"
	"os"
	"time"
)

func Select(c chan int, quits chan struct{}) {
	fmt.Println("excecuted2")
	for {
		time.Sleep(time.Second)
		select {
		case <-c:
			fmt.Println("Recieved")
		case <-quits:
			fmt.Println("Quits")
			//break
			os.Exit(1)
		}

	}

}

func main() {
	c := make(chan int)
	quits := make(chan struct{})
	go func() {
		c <- 1
		quits <- struct{}{}
		fmt.Println("excecuted1")
	}()
	go Select(c, quits)
	select {}

}
