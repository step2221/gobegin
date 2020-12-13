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

	fmt.Printf("Ксюше %d лет\n", ages["Ксюша"])
}
