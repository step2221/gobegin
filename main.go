package main

import (
	"fmt"
)

func main() {

	/*worker := scheduler.NewScheduler()
	fmt.Println(testmod.Hi("rob"))

	worker.Add(context.Background(), func(ctx context.Context) {
		fmt.Printf("Текущее время %s\n", time.Now())
	}, time.Second*1)
	time.Sleep(time.Minute)*/

	ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))
	spawnEmployees(ds)
	fmt.Println(ds.get(3))

	printType(2)
	printType("sdf")
	printType([]string{"1", "2"})
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(Employee{id: i})
	}
}

func printType(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("this is int")
	case string:
		fmt.Println("this is string")
	default:
		fmt.Println("this is not int and not string")
	}

	/*if _, ok := value.(string); ok {
		fmt.Println("this is string")
	} else if _, ok := value.(int); ok {
		fmt.Println("this is int")
	} else {
		fmt.Println("this is not int and not string")
	}
	*/
}
