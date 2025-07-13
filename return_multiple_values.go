package main


import "fmt"

func sayHello() (string, int) {
	return "Hello", 69
}

func main() {
	hello, number := sayHello()
	fmt.Println(hello, number)
}