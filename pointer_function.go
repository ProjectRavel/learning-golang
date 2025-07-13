package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeAddressToNegaraLain(address *Address) {
	address.Country = "Negara Lain"
	fmt.Println(address)
}

func main() {
	alamat := Address{"Depok", "Jawa Barat", ""}

	fmt.Println("Sebelum", alamat)

	ChangeAddressToNegaraLain(&alamat)
	fmt.Println("Setelah", alamat)

}
