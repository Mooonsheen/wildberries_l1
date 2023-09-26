// Дана структура Human (с произвольным набором полей и методов). Реализовать
// встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

package main

import "fmt"

// Родительская сруктура
type Human struct {
	name string
	age  int
}

func (h *Human) GetName() string {
	return h.name
}

func (h *Human) GetAge() int {
	return h.age
}

// Имплементация методов родительской структуры
type Action struct {
	Human
}

func (a *Action) SayHello() {
	fmt.Printf("Привет, меня зовут %s!\n", a.GetName())
}

func main() {
	human := Human{
		name: "Иван",
		age:  30,
	}

	action := Action{
		Human: human,
	}

	action.SayHello()
	fmt.Printf("Мне %d лет\n", action.GetAge())
}
