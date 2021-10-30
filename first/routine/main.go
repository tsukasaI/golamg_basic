package main

import (
	"fmt"
)

const NCPU = 4

func do(top, end int, data []int, c chan int) {
	for i := top; i < end; i++ {
		data[i] = data[i] * 5
	}
	c <- data[20]
}

func main() {
	s := make([]int, 50, 100)
	c := make(chan int, NCPU)

	for i := 0; i < len(s); i++ {
		s[i] = i
	}

	for i := 0; i < NCPU; i++ {
		go do(i*len(s)/NCPU, (i+1)*len(s)/NCPU, s, c)
	}
	for i := 0; i < NCPU; i++ {
		<-c
	}

	for _, i := range s {
		fmt.Printf("%d\n", i)
	}
}
