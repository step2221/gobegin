package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/robteix/testmod"
	"github.com/zhashkevych/scheduler"
)

type employee struct {
	id     int
	name   string
	sex    string //пол
	age    int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

type dumbStorage struct{}

func newDumbStorage() *dumbStorage {
	return &dumbStorage{}
}
func (s *dumbStorage) insert(e employee) error {
	fmt.Printf("вставка пользователя с ID: %d прошла успешно\n", e.id)
	return nil
}
func (s *dumbStorage) get(id int) (employee, error) {
	e := employee{id: id}
	return e, nil
}
func (s *dumbStorage) delete(id int) error {
	fmt.Printf("удаление пользователя с id : %d прошло успешно", id)
	return nil
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]employee),
	}
}

func (s *memoryStorage) insert(e employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) get(id int) (employee, error) {
	e, exists := s.data[id]
	if !exists {
		return employee{}, errors.New("такого сотрудника не существует")
	}
	return e, nil
}

func (s *memoryStorage) delete(id int) error {
	delete(s.data, id)
	return nil
}

func main() {

	worker := scheduler.NewScheduler()
	fmt.Println(testmod.Hi("rob"))

	worker.Add(context.Background(), func(ctx context.Context) {
		fmt.Printf("Текущее время %s\n", time.Now())
	}, time.Second*1)
	time.Sleep(time.Minute)

	/*ms := newMemoryStorage()
	ds := newDumbStorage()

	spawnEmployees(ms)
	fmt.Println(ms.get(3))
	spawnEmployees(ds)
	fmt.Println(ds.get(3))

	printType(2)
	printType("sdf")
	printType([]string{"1", "2"})*/
}

func spawnEmployees(s storage) {
	for i := 1; i <= 10; i++ {
		s.insert(employee{id: i})
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
