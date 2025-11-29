package stringfunctions

import (
	"fmt"
	"strings"
)

func Demo2() {
	isim := "Dogancan"
	fmt.Println(strings.HasPrefix(isim, "s"))
	fmt.Println(strings.HasSuffix(isim, "an"))
	fmt.Println(strings.Replace(isim, "a", "e", 2))
	fmt.Println(strings.Replace(isim, "a", "e", 1))
	fmt.Println(strings.Split(isim, "-"))
	fmt.Println(strings.Split(isim, "a"))
	fmt.Println(strings.Repeat(isim, 3))
}
