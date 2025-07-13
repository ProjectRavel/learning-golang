package main

import "fmt"


type Customer struct {
	Name, Address string
	Age int
}

func main(){
	customer := Customer{
		Name: "Rafael Pandu",
		Address: "Jl. Kebon Jeruk",
		Age: 20,
	}

	fmt.Println(customer.Name)
	fmt.Println(customer.Address)
	fmt.Println(customer.Age)
}