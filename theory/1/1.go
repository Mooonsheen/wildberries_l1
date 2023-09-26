// Какой самый эффективный способ конкатенации строк?

// В языке Go самый эффективный способ конкатенации строк - использование пакета `strings`
// и функции `Join`, которая объединяет срез строк в одну строку, используя заданный разделитель.

package main

import (
	"bytes"
	"strings"
	"testing"
)

const (
	countOfOperations = 100
)

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str string
		for j := 0; j < countOfOperations; j++ {
			str += "x"
		}
	}
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str bytes.Buffer
		for j := 0; j < countOfOperations; j++ {
			str.WriteString("x")
		}
	}
}

func BenchmarkStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var str strings.Builder
		for j := 0; j < countOfOperations; j++ {
			str.WriteString("x")
		}
	}
}
