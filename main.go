package main

import (
	"fmt"
	"time"

	"github.com/robteix/testmod"
	"github.com/step2221/gobegin/modules/storage"
)

func calculateSomething(e int) {
	t := time.Now()
	result := 0
	for i := 0; i <= e; i++ {
		result += i * 2
		time.Sleep(time.Millisecond * 3)
	}
	fmt.Printf("Результат: %d; Прошло времени %s", result, time.Since(t))
}

func main() {
	t := time.Now()
	fmt.Printf("Старт: %s\n", t.Format(time.RFC3339))

	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\r%c", r)
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()
	go calculateSomething(1000)

	go calculateSomething(2000)

	time.Sleep(time.Second * 8)
	fmt.Printf("Время выполнения программы: %s\n", time.Since(t))

	fmt.Println(testmod.Hi("rob"))
	/*worker := scheduler.NewScheduler()


	worker.Add(context.Background(), func(ctx context.Context) {
		fmt.Printf("Текущее время %s\n", time.Now())
	}, time.Second*1)
	time.Sleep(time.Minute)*/

	/*ms := storage.NewMemoryStorage()
	ds := storage.NewDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.Get(3))
	spawnEmployees(ds)
	fmt.Println(ds.Get(3))

	printType(2)
	printType("sdf")
	printType([]string{"1", "2"})*/
}

func spawnEmployees(s storage.Storage) {
	for i := 1; i <= 10; i++ {
		s.Insert(storage.Employee{
			Id: i,
		})
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
