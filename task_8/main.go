// Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import "fmt"

func setBit(n int64, pos uint, bitValue int) int64 {
	uwu := int64(1) << (pos - 1) // Создание маски с единицей на позиции нужного бита
	n &= ^uwu                    // Логическое бинарное И с инвертированной маской. Сбрасыват нужный бит до значения 0
	if bitValue == 1 {
		n |= uwu // При необходимых условиях меняет значение бита
	}
	return n
}

func main() {
	var num int64 // Исходное число
	fmt.Printf("Введите число: ")
	fmt.Scan(&num)
	var pos uint // Позиция бита, который мы хотим установить
	fmt.Printf("Введите номер бита: ")
	fmt.Scan(&pos)
	var bitValue int // Значение бита, которое мы хотим установить
	fmt.Printf("Введите желаемое значение бита: ")
	fmt.Scan(&bitValue)

	result := setBit(num, pos, bitValue)
	fmt.Printf("Результат: %d\n", result)
}
