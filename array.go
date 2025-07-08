package main

import "fmt"

func main() {
	var siswa [3]string

	siswa[0] = "Rafael"
	siswa[1] = "Cesya"
	siswa[2] = "Atila"

	fmt.Println(siswa)

	// atau bisa juga

	var angka = [...]int{90, 80, 70, 100}

	fmt.Println(len(angka))
}
