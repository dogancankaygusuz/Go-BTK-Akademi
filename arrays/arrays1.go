package arrays

import "fmt"

func DemoArrays() {
	var numbers [5]int

	fmt.Println(numbers)

	numbers[2] = 4
	numbers[4] = 2
	fmt.Println(numbers)

	var citys [3]string
	citys[0] = "ankara"
	citys[2] = "istanbul"

	for i := 0; i < 3; i++ {
		fmt.Println(citys[i])
	}
	fmt.Println(citys)
}
