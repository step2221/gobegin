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
	/*
		for i := 0; i <= 5; i++ {
			fmt.Println(i)
		}*/

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
	/*var arr [3]int
	fillArray(arr)
	fmt.Println(arr)

	brr := fillArraySuch(arr)
	fmt.Println(brr)*/

	// срез
	todoList1 := []string{"пункт 1", "пункт 2", "пункт 3"}
	fmt.Println("Длина списка:", len(todoList1))
	fmt.Println("Емкость списка:", cap(todoList1))

	todoList1 = append(todoList1, "Пунт 4 новый")

	fmt.Println("Длина списка после добавления:", len(todoList1))
	fmt.Println("Емкость списка после добавления:", cap(todoList1))

	newtodolist := append(todoList1, "Пункт 5", "Пункт 6")
	fmt.Println("Длина списка после добавления:", len(todoList1))
	fmt.Println("Емкость списка после добавления:", cap(todoList1))
	fmt.Println("Длина списка после добавления:", len(newtodolist))
	fmt.Println("Емкость списка после добавления:", cap(newtodolist))

	for _, item := range todoList1 {
		fmt.Printf("%s\n", item)
	}
	for _, item := range newtodolist {
		fmt.Printf("%s\n", item)
	}
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
