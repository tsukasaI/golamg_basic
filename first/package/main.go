package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	fmt.Printf("%v\n", len(args))
	for _, i := range args {
		fmt.Printf("%v\n", i)
	}

	err := os.Mkdir("dirname", 0666)
	if err != nil {
		fmt.Printf("%s\n", err.Error())
	}
	os.Remove("dirname")
	os.Exit(0)
	return
}
