package variables

import "fmt"

func Demo1() {
	var metin string = "Merhaba Dünya1"
	var tamsayi int = 22
	var ondalikli float32 = 34.5
	var durum bool = true
	dgncn := 24.3
	dgncn1 := "metinyazdim"

	fmt.Println(dgncn)
	fmt.Println(dgncn1)
	fmt.Printf("Veri Tipi: %T", dgncn)
	fmt.Println("Merhaba")
	fmt.Println(ondalikli)
	fmt.Print(tamsayi)
	fmt.Println(metin)
	fmt.Println(20 - 23*4)
	fmt.Println(durum)

	durum = metin == dgncn1
	fmt.Println(durum)

}
