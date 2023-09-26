// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

func intersection(nums1, nums2 []int) []int {
	set := make(map[int]bool)
	result := []int{}

	for _, num := range nums1 {
		set[num] = true
	}

	for _, num := range nums2 {
		if set[num] {
			result = append(result, num)
		}
	}

	return result
}

func main() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{4, 5, 6, 7, 8}

	res := intersection(nums1, nums2)
	fmt.Println(res) // Выводит [4, 5]
}
