package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomNumber генерирует случайное число с указанным количеством цифр.
func GenerateRandomNumber(length int) (int64, error) {
	// Вычисляем максимальное значение как 10^length - 1
	max := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil).Sub(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(length)), nil), big.NewInt(1))

	// Генерация случайного числа
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}

	return n.Int64(), nil
}
