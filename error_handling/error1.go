package errorhandling

import (
	"fmt"
	"os"
)

func Demo1() {
	f, err := os.Open("dosya11.txt")

	if err != nil {
		if pErr, ok := err.(*os.PathError); ok {
			fmt.Println("Dosya bulunamadı: ", pErr.Path)
			return
		} else {
			fmt.Println("Dosya bulunamadı")
			return
		}
	} else {
		fmt.Println(f.Name())
	}
}
