package inheritance

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	position  string
}

func (employee *Employee) GetName() string {

	fmt.Println("called GetEmployeeFUlleName")
	return employee.firstName + "," + employee.lastName
}

type SoleProprietor struct {
	employee Employee
}

func (employee Employee) SetName() {
	employee.firstName = "neelima"
}
