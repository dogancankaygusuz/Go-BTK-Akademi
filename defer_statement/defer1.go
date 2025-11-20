package deferstatement

import "fmt"

func A() {
	fmt.Println("A fonksiyonu çalıştı")
}

func B() {
	defer A() // B metotu çalıştıktan sonra A fonksiyonunu çalıştır
	defer C()
	fmt.Println("B fonksiyonu çalıştı")
}

func C() {
	fmt.Println("C fonksiyonu çalıştı")
}
