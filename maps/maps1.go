package maps

import "fmt"

func Demo1() {
	sozluk := make(map[string]string)

	sozluk["table"] = "masa"
	sozluk["pencil"] = "kalem"

	fmt.Println(sozluk)
	fmt.Println(sozluk["table"])

	deger, varMi := sozluk["lamb"]
	fmt.Println(deger)
	fmt.Println(varMi)

	sozluk1 := map[string]string{"1": "dogan", "2": "can"}

	fmt.Println(sozluk1)

}
