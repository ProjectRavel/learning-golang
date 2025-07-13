package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p1 := Person{
		Name: "Rafael Pandu",
		Age:  18,
	}

	p2 := &p1
	*p2 = Person{
		Name: "Rafael Pandu",
		Age:  18,
	}


	fmt.Println(p1)
	fmt.Println(p2)
}
