package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	address := new(Address)
	var addres2 *Address = address

	address.City = "Bandung"
	address.Province = "Jawa Barat"
	address.Country = "Indonesia"
	
	fmt.Println(address)
	fmt.Println(addres2)
}
