// Разработать программу нахождения расстояния между двумя точками, которые представлены
// в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

package main

import (
	"fmt"
	"math"
)

// Структура Point
type Point struct {
	x, y float64
}

// Конструктор Point
func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

// Метод для вычисления расстояния между двумя точками
func (p *Point) DistanceTo(other *Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создание двух точек
	p1 := NewPoint(1, 8)
	p2 := NewPoint(4, 6)

	// Вычисление расстояния между точками
	distance := p1.DistanceTo(p2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
