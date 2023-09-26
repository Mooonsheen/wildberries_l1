// Разработать программу, которая проверяет, что все символы в строке уникальные (
// true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

package main

import (
	"fmt"
)

func isUnique(str string) bool {
	// создаем пустую карту для хранения уникальных символов
	uniqueChars := make(map[rune]bool)

	// проходим по каждому символу в строке
	for _, char := range str {
		// проверяем, есть ли этот символ уже в карте
		if uniqueChars[char] {
			return false // символ уже присутствует - строка не уникальная
		}
		// добавляем символ в карту
		uniqueChars[char] = true
	}
	return true // все символы уникальны
}

func main() {
	str1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"

	fmt.Println(isUnique(str1)) // true
	fmt.Println(isUnique(str2)) // false
	fmt.Println(isUnique(str3)) // false
}
