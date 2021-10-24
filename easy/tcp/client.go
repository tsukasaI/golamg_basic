package main

import (
	"fmt"
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)
func main() {
	fmt.Println("input string. if you input 'quit' then quit this program")

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	} ()

	var sc = bufio.NewScanner(os.Stdin)
	var txt string
	for {
		if sc.Scan() {
			txt = sc.Text()
		}
		if txt == "quit" {
			break
		}
		txt += "\n"
		_, werr := conn.Write(([]byte)(txt))
		if werr != nil {
			log.Fatal(err)
		}
		time.Sleep(50 * time.Millisecond)
	}

	conn.Close()
	<-done
}
