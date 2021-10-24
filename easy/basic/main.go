package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var name string
	fmt.Printf("input string")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	name = scanner.Text()

	fmt.Printf("your input is %s", name)
}
