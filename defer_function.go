package main 

import "fmt"


func loggin(){
	fmt.Println("Selesai Memanggil")
}

func runApplication(){
	defer loggin()
	fmt.Println("Run Application")
}


func main(){
	runApplication()
}