package main

import (
	"fmt"
	"math"
)

type Value int

func (v Value) Twice() Value {
	return v * 2
}

type Point struct {X, Y float64}

func (p Point) Distance (q Point) float64 {
	d := math.Sqrt((p.X - q.X)*(p.X - q.X) + (p.Y - q.Y)*(p.Y - q.Y))
	return d
}


type Animal interface {
	Cry()
}

type Dog struct {}

type Cat struct {}


func (d *Dog) Cry() {
	fmt.Println("wow")
}

func (c *Cat) Cry() {
	fmt.Println("mew")
}

type Book struct {}

func Cry(any interface{}) {
	a, ok := any.(Animal)
	if !ok {
		fmt.Println("no cry")
		return
	}
	a.Cry()
}

func main() {
	var n, m Value

	n = 1
	m = n.Twice()

	fmt.Println(n, m)

	p1 := Point{5, 5}
	p2 := Point{2, 7}
	d := p1.Distance(p2)
	fmt.Println(d)

	dog := new(Dog)
	cat := new(Cat)

	Cry(dog)
	Cry(cat)
	book := new(Book)
	Cry(book)
}
