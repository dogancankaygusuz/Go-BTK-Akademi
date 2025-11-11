package loops

import "fmt"

func Demo1() {
	var str string = "dongu"

	i := 2

	for i <= 3 {
		fmt.Println(str)
		i++
	}

	for i := 3; i <= 5; i++ {
		fmt.Println("2.döngü çalişiyor")
		if i == 4 {
			fmt.Println("Şuanki sayi 4")
			fmt.Scanln(&i)
			fmt.Println("4 ün yeni değeri ", i)
		}
	}

}
