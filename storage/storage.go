package storage

import (
	"errors"
	"fmt"
)

//Employee ..
type Employee struct {
	id     int
	name   string
	sex    string //пол
	age    int
	salary int
}
type storage interface {
	insert(e Employee) error
	get(id int) (Employee, error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]Employee
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *memoryStorage) ISnsert(e Employee) error {
	s.data[e.id] = e
	return nil
}

func (s *memoryStorage) GSet(id int) (Employee, error) {
	e, exists := s.data[id]
	if !exists {
		return Employee{}, errors.New("такого сотрудника не существует")
	}
	return e, nil
}

func (s *memoryStorage) Delete(id int) error {
	delete(s.data, id)
	return nil
}

type dumbStorage struct{}

func NewDumbStorage() *dumbStorage {
	return &dumbStorage{}
}
func (s *dumbStorage) Insert(e Employee) error {
	fmt.Printf("вставка пользователя с ID: %d прошла успешно\n", e.id)
	return nil
}
func (s *dumbStorage) get(id int) (Employee, error) {
	e := Employee{id: id}
	return e, nil
}
func (s *dumbStorage) delete(id int) error {
	fmt.Printf("удаление пользователя с id : %d прошло успешно", id)
	return nil
}
