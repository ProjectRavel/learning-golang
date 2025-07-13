package main

import "fmt"

func random() any {
	return "random"
}

func main() {
	data := random()
	switch data.(type) {
	case string:
		fmt.Println("Data String")
	case int:
		fmt.Println("Data Integer")
	default:
		fmt.Println("Data Tidak Diketahui")
	}
}
