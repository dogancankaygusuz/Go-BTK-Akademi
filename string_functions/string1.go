package stringfunctions

import (
	"fmt"
	"strings"
)

func Demo1() {
	isim := "Dogancan"
	fmt.Println(strings.Count(isim, "a"))
	fmt.Println(strings.Contains(isim, "n"))
	fmt.Println(strings.Index(isim, "n"))
	fmt.Println(strings.ToUpper(isim))
	fmt.Println(strings.ToLower(isim))
}
