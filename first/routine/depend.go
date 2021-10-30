package main

import (
	"fmt"
)

func output(c chan string) {
	var str string
	for {
		str = <-c
		fmt.Printf("%s\n", str)
	}
	return
}

func input(c chan string, end chan int) {
	var str string

	for {
		fmt.Scanf("%s\n", &str)
		if str == "end" {
			end <- 1
		}
		c <- str
	}
	return
}

func main() {
	c := make(chan string)
	end := make(chan int)

	go output(c)
	go input(c, end)

	<-end

	fmt.Print("thank you")
}
