package main

import (
	"fmt"
)

func main() {
	var todoList = [3]string{
		"пункт 1", "пункт 2", "пункт 3",
	}
	fmt.Printf("Кол-во элементов в списке: %d \n", len(todoList))
	for index, item := range todoList {
		fmt.Printf("%d. %s \n", index, item)
	}
}
