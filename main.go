package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/channels"
	"golesson/conditionals"
	deferstatement "golesson/defer_statement"
	errorhandling "golesson/error_handling"
	"golesson/for_range"
	"golesson/functions"
	"golesson/interfaces"
	"golesson/loops"
	"golesson/maps"
	"golesson/pointers"
	"golesson/project"
	"golesson/restful"
	"golesson/slices"
	stringfunctions "golesson/string_functions"
	"golesson/structs"
	"golesson/variables"
)

func main() {
	variables.Demo1()
	conditionals.Demo1()
	loops.Demo1()
	arrays.DemoArrays()
	arrays.ArraysDemo2()
	arrays.ArraysDemo3()
	slices.SlicesDemo()
	fmt.Println(functions.Topla(5, 12))
	fmt.Println(functions.DortIslem(24, 3))

	sonuc1, sonuc2, sonuc3, sonuc4 := functions.DortIslem(25, 5)
	sonuc11, sonuc22, _, sonuc44 := functions.DortIslem(12, 3)
	fmt.Println(sonuc1)
	fmt.Println(sonuc2)
	fmt.Println(sonuc3)
	fmt.Println(sonuc4)
	fmt.Println(sonuc11)
	fmt.Println(sonuc22)
	fmt.Println(sonuc44)

	fmt.Println("VARIADIC")
	toplam := functions.ToplaVariadic(5, 6, 71, 5, 2, 94)
	fmt.Println(toplam)

	sayilar := []int{4, 52, 34, 14, 5, 12}
	fmt.Println(functions.ToplaVariadic(sayilar...))

	maps.Demo1()

	for_range.Demo1()

	p1 := 3
	pointers.Demo1(&p1)
	fmt.Println(p1)

	structs.Deneme()
	structs.Demo2()

	// go goroutines.CiftSayilar()
	// go goroutines.TekSayilar()

	ciftSayiCn := make(chan int)
	tekSayiCn := make(chan int)

	// Asenkron eşlenik zamanda çalışması için kullanıyoruz
	go channels.CiftSayilar(ciftSayiCn)
	go channels.TekSayilar(tekSayiCn)

	ciftSayiToplam, tekSayiToplam := <-ciftSayiCn, <-tekSayiCn
	carpim := ciftSayiToplam * tekSayiToplam
	fmt.Println("Carpim: ", carpim)

	interfaces.Demo1()

	deferstatement.B()
	deferstatement.Test()

	errorhandling.Demo1()

	interfaces.Demo3()

	errorhandling.Demo2()
	fmt.Println(errorhandling.TahminMetotu(3))

	stringfunctions.Demo1()
	stringfunctions.Demo2()

	restful.Demo1()
	restful.Demo2()

	project.GetAllProducts()
	project.AddProduct()
	project.GetAllProducts()
}
