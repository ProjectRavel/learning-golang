package main

import (
	"errors"
)

var ErrPembagianByZero = errors.New("pembagian by zero")

func Pembagian(nilai int, pembagian int) (int, error) {
	if pembagian == 0 {
		return 0, ErrPembagianByZero
	} else {
		result := nilai / pembagian
		return result, nil
	}
}

func main() {
	result, err := Pembagian(100, 4)
	if err == nil {
		println("Hasil", result)
	} else {
		println(err.Error())
	}
}
