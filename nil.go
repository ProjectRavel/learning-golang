package main

import "fmt"

// func contoh(name string) string {
// 	if name == "" {
// 		return nil
// 	} else {
// 		return name
// 	}
// }

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("")

	if data == nil {
		fmt.Println("Data Kosong")
	}else{
		fmt.Println(data)
	}
}
