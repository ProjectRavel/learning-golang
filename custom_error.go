package main

type validationError struct {
	message string
}

type notFoundError struct {
	message string
}

func (v *validationError) Error() string {
	return v.message
}

func (n *notFoundError) Error() string {
	return n.message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID Tidak Boleh Kosong"}
	}

	if id != "Rafael" {
		return &notFoundError{"Data Tidak Ditemukan"}
	}

	//ok 
	return nil
}

func main() {
	err := SaveData("Rafael", nil)

	if err != nil {
		println(err.Error())
	}
}
