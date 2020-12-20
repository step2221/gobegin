package storage

import (
	"errors"
	"fmt"
)

//Employee ..
type Employee struct {
	Id     int
	name   string
	sex    string //пол
	age    int
	salary int
}
type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type memoryStorage struct {
	data map[int]Employee
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data: make(map[int]Employee),
	}
}

func (s *memoryStorage) Insert(e Employee) error {
	s.data[e.Id] = e
	return nil
}

func (s *memoryStorage) Get(Id int) (Employee, error) {
	e, exists := s.data[Id]
	if !exists {
		return Employee{}, errors.New("такого сотрудника не существует")
	}
	return e, nil
}

func (s *memoryStorage) Delete(Id int) error {
	delete(s.data, Id)
	return nil
}

type dumbStorage struct{}

func NewDumbStorage() *dumbStorage {
	return &dumbStorage{}
}
func (s *dumbStorage) Insert(e Employee) error {
	fmt.Printf("вставка пользователя с ID: %d прошла успешно\n", e.Id)
	return nil
}
func (s *dumbStorage) Get(Id int) (Employee, error) {
	e := Employee{Id: Id}
	return e, nil
}
func (s *dumbStorage) Delete(Id int) error {
	fmt.Printf("удаление пользователя с id : %d прошло успешно", Id)
	return nil
}
