package errorhandling

import "fmt"

type BorderException struct {
	parameter int
	message   string
}

func (b *BorderException) Error() string {
	return fmt.Sprintf("%d---%s", b.parameter, b.message)
}

func TahminMetotu(tahmin int) (string, error) {
	if tahmin < 1 || tahmin > 100 {
		return "", &BorderException{parameter: tahmin, message: "Sınırların dışında kalındı."}
	}
	return "Bildiniz", nil
}
