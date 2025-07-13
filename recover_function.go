package main

import "fmt"

func endApp() {
	message := recover()

	if message != nil {
		fmt.Println("Error Message", message)
	}
	fmt.Println("End App")
}

func runApp(error bool) {
	defer endApp()
	if error == true {
		panic("ERROR")
	}

	fmt.Println("Run App")
}

func main() {
	runApp(true)
}
