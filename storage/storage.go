package storage

import (
	"errors"
	"fmt"
	"sync"
)

//Employee ..
type Employee struct {
	ID     int
	Name   string
	Sex    string //пол
	Age    int
	Salary int
}
type Storage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Update(id int, e Employee)
	Delete(id int) error
}

type memoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

func NewMemoryStorage() *memoryStorage {
	return &memoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

func (s *memoryStorage) Insert(e *Employee) error {
	s.Lock()
	e.ID = s.counter
	s.data[e.ID] = *e
	s.Unlock()
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
	fmt.Printf("вставка пользователя с ID: %d прошла успешно\n", e.ID)
	return nil
}
func (s *dumbStorage) Get(Id int) (Employee, error) {
	e := Employee{ID: Id}
	return e, nil
}
func (s *dumbStorage) Delete(Id int) error {
	fmt.Printf("удаление пользователя с id : %d прошло успешно", Id)
	return nil
}
