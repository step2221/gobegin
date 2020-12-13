package main

import "fmt"

func main() {
	/*ages := make(map[string]int)
	ages["Ксюша"] = 7
	ages["Андрей"] = 5
	ages["Саша"] = 8

	for key, value := range ages {
		fmt.Printf("%s - %d\n", key, value)
	}*/
	ages := map[string]int{
		"Ксюша":  7,
		"Андрей": 5,
		"Саша":   8,
	}

	delete(ages, "Саша")
	// Мапы как и срезы являются ссылкой

	fmt.Printf("Ксюше %d лет\n", ages["Ксюша"])

	age, exists := ages["Антон"]
	if !exists {
		fmt.Println("Антона нет в списке")
	} else {
		fmt.Printf("Антона %d лет\n", age)
	}
}
