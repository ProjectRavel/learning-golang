package main

import "fmt"

func main(){
	// saat kita membuat variable maka wajib menyebutkan tiper data variabel tersebut
	// namun jika kita langsung menginisialisasikan data pada variable nya. maka kita tidak wajib menyebutkan tipe datanya
	var name string

	name = "Rafael Pandu"
	fmt.Println(name)

	name = "ProjectRavel"
	fmt.Println(name)

	var tanpaTipeData = "hai"
	fmt.Println(tanpaTipeData)

	// := ini sama aja seperti membuat var tanpa var langsung 	
	iniTanpaVar := "test"
	fmt.Println(iniTanpaVar)


	// multiple var
	var (
		firstName = "Rafael"
		middleName = "Pandu"
		lastName = "Sumanti"
	)

	fmt.Println(firstName, middleName, lastName)
}