package main

import (
	"fmt"
	"learning-golang/database"
	_ "learning-golang/internal"
)

func main() {
	fmt.Println(database.GetDataBase())
	
}
