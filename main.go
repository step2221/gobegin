package main

import (
	"fmt"
)

func main() {
	/*var todoList = [3]string{
		"пункт 1", "пункт 2", "пункт 3",
	}*/
	//Либо так ...
	/*var todoList [3]string
	todoList[0] = "пункт 1"*/
	//Либо так ...
	var todoList = [...]string{"пункт 1", "пункт 2", "пункт 3"}

	fmt.Printf("Кол-во элементов в списке: %d \n", len(todoList))
	/*for index, item := range todoList {
		fmt.Printf("%d. %s \n", index, item)
	}*/
	for _, item := range todoList {
		fmt.Printf("%s\n", item)
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	//Так бесконечный цикл
	/*for i := 0; i <= 5; {
		fmt.Println(i)
	}*/

	//Так цикл до завершения программы
	/*for i := 0; ; i++ {
		fmt.Println(i)
	}*/

	//Бесконечный цикл
	/*i := 0
	for {
		if i == 10 {
			break
		}
		fmt.Println(i)
		i++

	}*/
	var arr [3]int
	fillArray(arr)
	fmt.Println(arr)

	brr := fillArraySuch(arr)
	fmt.Println(brr)

}

//fillArray .
func fillArray(arr [3]int) {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}

	fmt.Println("fillArray():", arr)
}

//fillArraySuch .
func fillArraySuch(arr [3]int) [3]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = i
	}
	return arr

}
