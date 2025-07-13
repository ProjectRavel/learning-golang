package main

import "fmt"

func main() {
	// counter := 1

	// for counter <= 10 {
	// 	fmt.Println("Perulangan ke",counter)
	// 	counter++
	// }


	for counter := 0; counter <= 10000; counter++ {
		fmt.Println("Counter ke", counter)
	}


	siswa := []string{"Rafael", "Zeta", "kobo"}

	for index, name := range siswa {
		fmt.Println(index, name)
	}

	for _,name := range siswa {
		fmt.Println(name)
	}
}