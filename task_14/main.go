// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
)

// Функция определения типа переменной
func typeof(i interface{}) string {
	switch i.(type) {
	case string:
		return "String"
	case int:
		return "Number"
	case bool:
		return "Bool"
	case chan interface{}:
		return "Channel"
	default:
		return "Other"
	}
}

func main() {
	var val interface{} = false // Переменная типа interface{}
	fmt.Printf("%s", typeof(val))
}
