// Что такое интерфейсы, как они применяются в Go?

// В Go интерфейсы представляют собой набор методов, которые определяют поведение объекта.
// Они определяют, какие методы должны быть реализованы в типе данных, чтобы этот тип
// данных удовлетворял интерфейсу.

// Использование интерфейсов позволяет достичь полиморфизма в Go. Это означает, что мы
// можем использовать переменные интерфейсного типа для работы с различными типами данных,
// которые реализуют этот интерфейс.

package main

// import (
// 	"fmt"
// )

// type Stringer interface { // объявляем интерфейс с методом String()
// 	String()
// }

// type Person struct { // объявляем стуркуру, у которой будет определять метод String()
// 	name string // для реализации интефейса Stringer
// }

// func (p Person) String() string { // Реализуем у структуры Person метод String(). Таким
// 	return "My name is " + p.name // образом Person реализует интерфейс Stringer
// }

// func main() {
// 	p := Person{"Nikita"}
// 	fmt.Print(Person.String(p))
// }
