package datastructures

import (
	"fmt"
)

type employee struct {
	name           string
	address        string
	yearsOfService int
}

func (emp *employee) getEmployeeName() string {
	return emp.name
}

func createMultipleEmployees() {

	employees := []employee{}
	for i := 0; i < 100; i++ {
		elem := employee{"name", "address", 10}
		employees = append(employees, elem)
	}

}

func CreateAndIterateMap() map[string]int {
	fmt.Println("in datastructures.go ")
	m := make(map[string]int)

	m["a"] = 1
	m["b"] = 1
	for key, val := range m {
		fmt.Print("keys is " + key)
		fmt.Printf("value is %d\n", val)
	}
	return m
}

func LoopOverArray(inputArray []int) {

	for index, val := range inputArray {

		fmt.Printf("Index is %d Element is %d ", index, val)

	}
}
