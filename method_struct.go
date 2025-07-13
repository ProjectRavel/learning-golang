package main

import "fmt"

type Person struct {
	first string
	last  string
	age   int
}

func (p Person) speak() {
	fmt.Println("Hello, my name is", p.first, p.last, "and I am", p.age, "years old.")
}

func main() {
	p1 := Person{
		first: "Rafael",
		last:  "Pandu",
		age:   18}
	p1.speak()

	fmt.Println(p1)
}
