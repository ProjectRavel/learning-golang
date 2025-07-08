package main

import "fmt"

func main() {
	siswa := [...]string{"Rafael", "Cesya", "Atila", "Fathir", "GibraN"}
	slice1 := siswa[2:5]
	slice2 := siswa[:5]
	slice3 := siswa[2:]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jum'at", "sabtu", "minggu"}
	daysSlice1 := days[5:7]
	daysSlice1[0] = "Baru"

	daysSlice2 := append(daysSlice1, "libur")

	iniArray := [...]int{1,2,3}
	iniSlice := []int{1,2,3}

	
	fmt.Println(daysSlice1)
	fmt.Println(days)
	fmt.Println(daysSlice2)
	fmt.Println(cap(daysSlice1))

}