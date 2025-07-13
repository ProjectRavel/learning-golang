package main

import "fmt"

func main() {
	name := "Zea"

	if name == "Zetadasdsadsdas" {
		fmt.Println("HALOOO SAYANG AKU CINTA AKUUUUUU")
	} else if name == "Kobo" {
		fmt.Println("HALO AKU SAYANG KAMUUU")
	} else {
		fmt.Println("who r u bro")
	}


	// short statement
	if length := len(name); length < 5 {
		fmt.Println("Nama anda sangat pendek sekali")
	} else{
		fmt.Println("Nama anda sangat panjang sekali")
	}
}
