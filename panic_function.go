package main

import "fmt"

func endApp() {
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error == true {
		panic("ERROR")
	}
	fmt.Println("Run App")
}

func main(){
	runApp(true)
}
