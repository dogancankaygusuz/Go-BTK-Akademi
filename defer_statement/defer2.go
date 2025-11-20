package deferstatement

import "fmt"

type product struct {
	productName string
	unitPrice   int
}

func (p product) Save() {
	fmt.Println("Ürün Kaydedildi: ", p.productName)
}

func Test() {
	p := product{productName: "Laptop", unitPrice: 500}
	defer p.Save()
	fmt.Println("İşlem başarılı")
}
