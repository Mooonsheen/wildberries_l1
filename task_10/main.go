// Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
// Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.

package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Сортируем температуры по возрастанию
	sort.Float64s(temperatures)

	groups := make(map[float64][]float64)

	// Разбиваем температуры на группы с шагом в 10 градусов
	for _, temp := range temperatures {
		key := math.Floor(temp/10) * 10
		groups[key] = append(groups[key], temp)
	}

	// Выводим результат
	for key, group := range groups {
		fmt.Println(key, ":", group)
	}
}
