package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fname := `sampleioutli,dat`

	txt := "this is a writing sample"

	werr := ioutil.WriteFile(fname, ([]byte)(txt), 0666)
	if werr != nil {
		_ = fmt.Errorf("%sに書き込みできません", fname)
	}

	data, rerr := ioutil.ReadFile(fname)
	if rerr != nil {
		_ = fmt.Errorf("%sを読み込めません", fname)
	} else {
		fmt.Print(string(data))
	}
}
