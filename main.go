package main

import (
	"fmt"
	"golesson/arrays"
	"golesson/conditionals"
	"golesson/functions"
	"golesson/loops"
	"golesson/slices"
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
}
