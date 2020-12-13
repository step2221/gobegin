package main

import (
	"errors"
	"fmt"
)

const pi float32 = 3.1415

func main() {

	PrintCircleArea(-2)

}

//CalculateCircleArea ..Рассчитывает площадь окружности
func CalculateCircleArea(radius int) (float32, error) {
	if radius <= 0 {
		return float32(0), errors.New("радиус круга не может быть отрицательным")
	}
	return pi * (float32(radius) * float32(radius)), nil
}

//PrintCircleArea Вывод в терминал
func PrintCircleArea(radius int) {

	circleArea, err := CalculateCircleArea(radius)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Радиус круга: %d см. \n", radius)
	fmt.Printf("Формула для расчета круга: A=пr2\n")
	fmt.Printf("Площадь круга: %f32 см. кв.\n", circleArea)
}
