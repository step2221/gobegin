package main

import "fmt"

const pi float64 = 3.1415

func main() {
	PrintCircleArea(2)
	PrintCircleArea(4)

	fmt.Println("Площадь круга c радиусом 5см=", CalculateCircleArea(5))
}

//CalculateCircleArea ..Рассчитывает площадь окружности
func CalculateCircleArea(radius int) float64 {
	return pi * (float64(radius) * float64(radius))
}

//PrintCircleArea Вывод в терминал
func PrintCircleArea(radius int) {
	fmt.Printf("Радиус круга: %d см. \n", radius)
	fmt.Printf("Формула для расчета круга: A=пr2\n")

	circleArea := CalculateCircleArea(radius)
	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleArea)
}
