package storage

import (
	"errors"
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

type MemoryStorage struct {
	counter int
	data    map[int]Employee
	sync.Mutex
}

//NewMemoryStorage ..
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data:    make(map[int]Employee),
		counter: 1,
	}
}

//Insert .
func (s *MemoryStorage) Insert(e *Employee) error {
	s.Lock()
	e.ID = s.counter
	s.data[e.ID] = *e
	s.counter++
	s.Unlock()
	return nil
}

//Get ..
func (s *MemoryStorage) Get(Id int) (Employee, error) {
	s.Lock()
	defer s.Unlock()

	employee, ok := s.data[Id]
	if !ok {
		return Employee{}, errors.New("такого сотрудника не существует")
	}
	return employee, nil
}

//Update ..
func (s *MemoryStorage) Update(id int, e Employee) {
	s.Lock()
	s.data[id] = e
	s.Unlock()
}

//Delete .
func (s *MemoryStorage) Delete(Id int) {
	s.Lock()
	delete(s.data, Id)
	s.Unlock()

}
