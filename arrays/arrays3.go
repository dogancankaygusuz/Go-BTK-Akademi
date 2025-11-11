package arrays

import "fmt"

func ArraysDemo3() {
	var numbers [3][2]int

	numbers[0][0] = 7
	numbers[0][1] = 5
	numbers[1][1] = 4
	numbers[2][0] = 6

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 1; j++ {
			fmt.Print(numbers[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
