// Разработать программу, которая перемножает, делит, складывает, вычитает две
// числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	// Установка значений переменных
	a.SetString("12345678901234567890", 10)
	b.SetString("98765432109876543210", 10)

	// Перемножение
	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Println("Умножение:", mul)

	// Деление
	div := new(big.Int)
	div.Div(a, b)
	fmt.Println("Деление:", div)

	// Сложение
	add := new(big.Int)
	add.Add(a, b)
	fmt.Println("Сложение:", add)

	// Вычитание
	sub := new(big.Int)
	sub.Sub(a, b)
	fmt.Println("Вычитание:", sub)
}
