package main

import(
	"fmt"
)

type myError struct {
	Code int
	Msg string
}

func (e myError) Error() string {
	return fmt.Sprintf("Code: %d, Msg: %s", e.Code, e.Msg)
}

func prosesData(x int) error {
	if x < 0 {
		return &myError{
			Code: 404,
			Msg: "Data Tidak Ditemukan",
		}
	}

	return nil
}

func main(){
	err := prosesData(-1)
	if err != nil {
		fmt.Println("TERJADI ERROR", err.Error())
	} else {
		fmt.Println("BERHASIL")
	}
}