package structs

import "fmt"

func Deneme() {
	fmt.Println(product{"Bilgisayar", 5000, "X"})
	fmt.Println(product{name: "Telefon", brand: "Y"})
}

type product struct {
	name      string
	unitPrice float64
	brand     string
}
