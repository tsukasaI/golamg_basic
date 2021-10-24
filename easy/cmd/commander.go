package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("引数を２個指定してください。")
		os.Exit(1)
	}

	var x, y int
	x, _ = strconv.Atoi(os.Args[1])
	y, _ = strconv.Atoi(os.Args[2])
	fmt.Printf("実行ファイル名: %s\n", os.Args[0])
	fmt.Printf("引数1: %s\n", os.Args[1])
	fmt.Printf("引数2: %s\n", os.Args[2])
	fmt.Printf("%d + %d = %d\n", x, y, x + y)
}
