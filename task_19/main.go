// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var str string
	fmt.Scan(&str)
	reversed := reverseString(str)
	fmt.Println(reversed)
}

func reverseString(str string) string {
	reversed := make([]byte, 0, len(str))
	for len(str) > 0 {
		r, size := utf8.DecodeLastRuneInString(str)
		reversed = append(reversed, []byte(string(r))...)
		str = str[:len(str)-size]
	}
	return string(reversed)
}
