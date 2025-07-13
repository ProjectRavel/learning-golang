package main

import "fmt"

func main() {
	person := map[string]string{
		"name":  "Rafael",
		"class": "XI RPL 1",
	}

	fmt.Println(person["name"])
	fmt.Println(person["class"])
	fmt.Println(person)
	delete(person, "class")

	fmt.Println(person)
}