// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x)
// из массива, во второй — результат операции x*2, после чего данные из
// второго канала должны выводиться в stdout.

package main

import (
	"fmt"
)

func main() {
	// Создаем два канала
	input := make(chan int)
	output := make(chan int)

	// Запускаем горутину для записи чисел из массива в канал input
	go func() {
		numbers := []int{1, 2, 3, 4, 5}

		for _, num := range numbers {
			input <- num
		}

		close(input)
	}()

	// Запускаем горутину для выполнения операции x*2 и записи результата в канал output
	go func() {
		for num := range input {
			result := num * 2
			output <- result
		}

		close(output)
	}()

	// Выводим числа из канала output в stdout
	for result := range output {
		fmt.Println(result)
	}

}
