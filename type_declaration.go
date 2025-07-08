package main

import "fmt"

func main(){
	type noKtp string

	var ktp = 69

	fmt.Println(noKtp(ktp))
}