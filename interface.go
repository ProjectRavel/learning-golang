package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName){	
	println("Hello:", value.GetName())
}

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main(){
	rafael := Person{Name: "Rafael Pandu"}
	fox := Animal{Name: "Fox"}

	fmt.Println(fox.GetName())
	SayHello(rafael) 
}