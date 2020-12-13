package main

import "fmt"

type employee struct {
	name   string
	sex    string //пол
	age    int
	salary int
}

//Кастомный тип
type age int

func (a age) isAdult() bool { //метод с ресивером
	return a >= 18
}

func main() {
	/*ages := make(map[string]int)
	ages["Ксюша"] = 7
	ages["Андрей"] = 5
	ages["Саша"] = 8

	for key, value := range ages {
		fmt.Printf("%s - %d\n", key, value)
	}*/
	myAge := age(20)
	fmt.Println("Я совершеннолетний?", myAge.isAdult())

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

	employee1 := newEmployee("Вася", "М", 25, 1500)
	employee2 := newEmployee("Петя", "М", 28, 2000)
	setName(&employee1, "Геннадий") //передаем  указатель
	employee1.setNameEm("Конечно Вася")

	fmt.Printf("%v\n", employee1.getInfo())

	fmt.Printf("%v\n", employee2.getInfo())
	//fmt.Printf("%v\n", employee2)
}

func newEmployee(name, sex string, age, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник: %s\n Возраст: %d\n Зарплата: %d\n", e.name, e.age, e.salary)
}

func setName(e *employee, name string) { //функция в которую передается указатель
	//на структуру, для изменеия самой структуры
	e.name = name
}
func (e *employee) setNameEm(name string) { //функция указатель в роли ресивера
	e.name = name
}
