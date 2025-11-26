package interfaces

import "fmt"

func isCheck(i interface{}) {
	sayi, ok := i.(int)
	fmt.Println(sayi, ok)
}

func Demo3() {
	isCheck(6)
	isCheck("dgncn")

	var x interface{} = "dogancan"
	isCheck(x)
}
