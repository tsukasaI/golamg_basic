package main

import "fmt"

const NCPU = 4

func handler(req chan int) {
	for {
		data := <-req

		for i := 2; i < data; i++ {
			if data%i == 0 {
				fmt.Printf("negative\n")
				break
			}
		}
	}
}

func main() {
	var data int
	req := make(chan int, 10)
	for i := 0; i < NCPU; i++ {
		go handler(req)
	}

	for {
		fmt.Scanf("%d", &data)
		req <- data
	}
}
