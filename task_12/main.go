// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import "fmt"

func main() {
	sequence := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]bool)
	for _, val := range sequence {
		set[val] = true
	}

	fmt.Println(set)
}
