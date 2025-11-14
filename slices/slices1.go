package slices

import "fmt"

func SlicesDemo() {
	names := make([]string, 3)
	fmt.Println(names)

	names[1] = "dogan"
	names[2] = "can"
	names = append(names, "kaygusuz")
	fmt.Println(names)

	newNames := make([]string, len(names))
	copy(newNames, names)
	fmt.Println(newNames, "size:", len(newNames))
	fmt.Println(names[1:3])
}
