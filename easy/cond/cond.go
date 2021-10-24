package main

import "fmt"

func main() {
	x := 0.0
	for i := 1; x < 2000; i++ {
		a := 40.0
		p := 1.04

		x = (x + a) * p
		fmt.Printf("%d年目: %.2f 万円\n", i, x)
	}

	for i, v := range []string{"apple", "facebook", "amazon"} {
		fmt.Println(i, v)
	}
}
