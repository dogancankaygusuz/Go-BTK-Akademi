package arrays

import (
	"fmt"
)

func ArraysDemo2() {
	numbers := [4]int{2, 4, 87, 4}
	maximumNumber := numbers[0]

	for i := 0; i < len(numbers); i++ {
		if maximumNumber < numbers[i] {
			maximumNumber = numbers[i]
		}
	}
	fmt.Println("Max Sayi: ", maximumNumber)
	fmt.Println(numbers)
}
