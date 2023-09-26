// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	text, _ := in.ReadString('\n')
	words := strings.Fields(text)

	// Проходимся по слайсу слов, применяя реверс-функцию к каждому слову
	for i := 0; i < len(words); i++ {
		words[i] = reverseString(words[i])
	}

	reversed := strings.Join(words, " ")
	fmt.Fprintf(out, "%s", reversed)
	out.Flush()
}

// Функция реверса слова
func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
