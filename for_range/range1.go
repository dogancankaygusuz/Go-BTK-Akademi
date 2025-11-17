package for_range

import "fmt"

func Demo1() {
	sehirler := []string{"istanbul", "ankara", "izmir"}

	for _, sehir := range sehirler {
		fmt.Println(sehir)
	}

	for i, sehir := range sehirler {
		fmt.Print(i, "-")
		fmt.Println(sehir)
	}

	sozluk := map[string]string{"book": "kitap", "table": "masa"}

	for k, v := range sozluk {
		fmt.Print(k, " : ", v, "\n")
	}
}
