package main

import (
	"fmt"
	"bufio"
	"log"
	"net"
	"strings"
	"time"
)

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		msg := input.Text()
		go func() {
			if strings.ToLower(msg) == "hello" {
				if time.Now().Hour() < 12 {
					fmt.Fprintln(c, "[good morning]")
				} else {
					fmt.Fprintln(c, "[good evening]")
				}
			} else {
				fmt.Fprintln(c, "[", strings.ToUpper(msg), "]")
			}
			time.Sleep(500 * time.Millisecond)
		} ()
	}
	c.Close()
}

func main() {
	fmt.Println("start server")

	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Print(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
