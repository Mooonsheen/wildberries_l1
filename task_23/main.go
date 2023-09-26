// Удалить i-ый элемент из слайса.

// Способ 1
// package main

// import "fmt"

// func main() {
// 	slc := []int{1, 2, 3, 4, 5}
// 	var i int
// 	fmt.Printf("Порядковый номер удаляемого элемента: ")
// 	fmt.Scan(&i)

// 	slc = append(slc[:i-1], slc[i:]...)

// 	fmt.Println("Слайс после удаления элемента:", slc)
// }

// Способ 2
package main

import "fmt"

func main() {
	// Исходный слайс
	slc := []int{1, 2, 3, 4, 5}
	var i int
	fmt.Printf("Порядковый номер удаляемого элемента: ")
	fmt.Scan(&i)
	copy(slc[i-1:], slc[i:])
	slc = slc[:len(slc)-1]

	fmt.Println("Слайс после удаления элемента:", slc)
}
