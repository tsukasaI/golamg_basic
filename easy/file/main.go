package main

import (
	"fmt"
	"os"
)

func main() {
	fname := `sample.dat`
	var file *os.File
	var err error

	file, err = os.Create(fname)
	if err != nil {
		_ = fmt.Errorf("%sを開けません。\n", fname)
	}

	txt := "this is an example"

	file.Write(([]byte)(txt))

	file.Close()

	const BUFSIZE = 1024

	file, err = os.Open(fname)
	if err != nil {
		_ = fmt.Errorf("%sを開けません。\n", fname)
	}
	defer file.Close()

	buf := make([]byte, BUFSIZE)

	for {
		n, err := file.Read(buf)
		if n == 0 {
			break
		}
		if err != nil {
			_ = fmt.Errorf("%sの読み込みでエラーが発生しました\n", fname)
			break
		}
		fmt.Print(string(buf[:n]))
	}
}
