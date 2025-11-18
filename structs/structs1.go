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

type customer struct {
	firstName string
	lastName  string
	age       int
}

func (c customer) save() {
	fmt.Println("Eklendi", c.firstName)
}

func Demo2() {
	c := customer{firstName: "dogancan", lastName: "kaygusuz", age: 22}
	c.save()
}
