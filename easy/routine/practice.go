package main

import (
	"fmt"
	"strings"
	"sync"
)
	var wg sync.WaitGroup

func convertToCapital(s string) {
	defer wg.Done()
	fmt.Printf("%s\n", strings.ToUpper(s))
}

func main() {
	wg.Add(1)
	var s string

	fmt.Println("input word")
	fmt.Scan(&s)

	go convertToCapital(s)

	wg.Wait()

}
