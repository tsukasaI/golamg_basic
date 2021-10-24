package main

import (
	"fmt"
	"time"
	"sync"
)

// var wg sync.WaitGroup

// func printjob(s string) {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		time.Sleep(10 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	wg.Add(3)
// 	fmt.Println("A started")
// 	go printjob("A")
// 	fmt.Println("B started")
// 	go printjob("B")
// 	fmt.Println("C started")
// 	go printjob("C")

// 	wg.Wait()
// 	fmt.Printf("\nend\n")
// }

// import (
// 	"fmt"
// 	"time"
// )
// func main() {
// 	done := make(chan bool)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(10 * time.Millisecond)
// 			fmt.Println("+")
// 		}
// 		done <- true
// 	}()

// 	for i := 0; i < 10; i++ {
// 		time.Sleep(10 * time.Millisecond)
// 		fmt.Println("/")
// 	}
// 	<-done

// 	fmt.Println("done")

// }

// func main () {
// 	ch1 := make(chan rune)
// 	ch2 := make(chan int)
// 	done := make(chan bool)

// 	go func() {
// 		s := "ABCDEFG"
// 		for _, c := range s {
// 			time.Sleep(10 * time.Millisecond)
// 			fmt.Printf("ch1から送信: %c\n", c)
// 			ch1 <- c
// 		}
// 		done <- true
// 	} ()

// 	go func () {
// 		for i := 0; i < 8; i++ {
// 			time.Sleep(8 * time.Millisecond)
// 			fmt.Printf("ch2から送信: %d\n", i)
// 			ch2 <- i + 1
// 		}
// 		done <- true
// 	} ()

// 	defer fmt.Println("end")
// 	count := 0

// 	for {
// 		select {
// 		case r := <-ch1:
// 			fmt.Printf("ch1から受信: %c\n", r)
// 		case n := <-ch2:
// 			fmt.Printf("ch1から受信: %d\n", n)
// 		case <-done:
// 			count++
// 			if count > 1 {
// 				return
// 			}
// 		}
// 	}
// }

// func main() {
// 	done := make(chan bool, 3)

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(10 * time.Millisecond)
// 			fmt.Println("+")
// 		}
// 		done <- true
// 	} ()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(9 * time.Millisecond)
// 			fmt.Println("-")
// 		}
// 		done <- true
// 	} ()

// 	go func() {
// 		for i := 0; i < 10; i++ {
// 			time.Sleep(8 * time.Millisecond)
// 			fmt.Println("\\")
// 		}
// 		done <- true
// 	} ()

// 	for i := 0; i < 10; i++ {
// 		fmt.Println("/")
// 		time.Sleep(5 * time.Millisecond)
// 	}

// 	for i := 0; i < 3; i++ {
// 		<-done
// 	}


// }

var wg sync.WaitGroup
var a int
var mu sync.Mutex

func jobPlus() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		t := a
		a = a + 1
		time.Sleep(10 * time.Millisecond)
		fmt.Printf("plus %d -> %d\n", t, a)
		mu.Unlock()
	}
}

func jobMinus() {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		t := a
		a = a - 1
		time.Sleep(11 * time.Millisecond)
		fmt.Printf("minus %d -> %d\n", t, a)
		mu.Unlock()
	}
}

func main() {
	a = 10
	wg.Add(2)

	go jobPlus()
	go jobMinus()

	wg.Wait()
}
