// К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
// Приведите корректный пример реализации.

// var justString string
// func someFunc() {
//   v := createHugeString(1 << 10)
//   justString = v[:100]
// }

// func main() {
//   someFunc()
// }

package main

import (
	"fmt"
	"strings"
)

func createHugeString(a int) string {
	str := "0"
	res := strings.Repeat(str, a)
	return res
}

func someFunc() {
	v := createHugeString(1 << 10)
	justString := v[:100]
	fmt.Print(justString)
}

func main() {
	someFunc()
}
