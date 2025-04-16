package main

import "fmt"

type person struct {
	name string
	age  int
}

type Employee struct {
	person
	empId  string
	salary float64
}

func (p person) introduce() {
	fmt.Printf("Hi,I'm %s and I'm %d years old.\n", p.name, p.age)
}

func (e Employee) introduce() {
	fmt.Printf("Hi,I'm %s, employee Id: %s,and I earn: %.2f\n", e.name, e.empId, e.salary)
}

func main() {
	employee := Employee{
		person: person{
			name: "John",
			age:  27,
		},
		empId:  "E001",
		salary: 50000,
	}
	fmt.Println("Name:", employee.name)
	fmt.Println("Age:", employee.age)
	fmt.Println("EmpId:", employee.empId)
	fmt.Println("Salary:", employee.salary)

	employee.introduce()
	employee.person.introduce()

}
