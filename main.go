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
}
