package main

import "fmt"

func main()  {
	// var a [5]int = [5]int{1, 3, 5, 7, 9}

	// n := a[4]
	// fmt.Println(n)

	// for i := 0; i < 5; i++ {
	// 	fmt.Printf("a[%d] = %d\n", i, a[i])
	// }

	// for i, v := range a {
	// 	fmt.Printf("a[%d] = %d\n", i, v)
	// }

	// var c = make([]int, 5, 20)
	// fmt.Println(len(c), cap(c))

	var a = [...]int{1,3,5,7,9}

	var s = a[1:4]

	var d = make([]int, 5, 20)

	copy(d, s)

	d = append(d, 21)

	fmt.Println(a)
	fmt.Println(s)
	fmt.Println(d)

	var m = map[string]int{}

	m["po"] = 5
	m["tv"] = 3
	fmt.Println(m)

	// struct
	type member struct {
		id string
		name string
		age int
	}

	var mp = new(member)

	*mp = member{"c9bjf3", "yayoi", 31}

	fmt.Println(*mp)

	var members = []member{}
	members = append(members, member{"fjiov", "jfigod", 31})
	fmt.Println(members)
}
