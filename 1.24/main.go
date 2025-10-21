package main

import (
	"fmt"
	"math"
)

// Point структура с приватными полями
type Point struct {
	x float64
	y float64
}

// NewPoint конструктор для создания новой точки
func NewPoint(x, y float64) Point {
	return Point{
		x: x,
		y: y,
	}
}

// Distance метод вычисления расстояния между двумя точками
func (p *Point) Distance(other Point) float64 {
	dx := p.x - other.x
	dy := p.y - other.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаём две точки
	p1 := NewPoint(3, 4)
	p2 := NewPoint(4, 4)

	// Вычисляем расстояние
	distance := p1.Distance(p2)

	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
