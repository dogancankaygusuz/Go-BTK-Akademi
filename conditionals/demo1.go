package conditionals

import "fmt"

func Demo1() {
	var hesap float32 = 1000
	var cekilmekIstenen float32 = 900
	if cekilmekIstenen > hesap {
		fmt.Println("Hesaptaki Para Yetersiz :(")
	} else {
		fmt.Println("Paranız Çekildi. " + fmt.Sprintf("%v", hesap-cekilmekIstenen))
	}
}
