package main

import "fmt"

type Person struct {
	Name string
	age  int
}

func (p *Person) married() {
	p.Name = "MR." + p.Name
}

func main() {
	rafael := Person{Name: "Rafael Pandu", age: 20}
	rafael.married()
	fmt.Println(rafael)
}
