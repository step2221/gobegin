package main

import (
	"fmt"
)

func main() {
	ages := make(map[string]int)
	ages["Ксюша"] = 7
	ages["Андрей"] = 5
	ages["Саша"] = 8

	for key, value := range ages {
		fmt.Printf("%s - %d\n", key, value)
	}
}
