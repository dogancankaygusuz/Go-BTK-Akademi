package errorhandling

import (
	"errors"
	"fmt"
)

func tahminEt(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", errors.New("1-100 arasında bir sayı giriniz.")
	}
	return "bildiniz", nil
}

func Demo2() {
	mesaj, _ := tahminEt(45)
	fmt.Println(mesaj)
}
