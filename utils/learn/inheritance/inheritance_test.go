package inheritance

import (
	"fmt"
	"testing"
)

func TestProprietor(t *testing.T) {
	fmt.Println("In test Propriter t")
	//proprietor := SoleProprietor{Employee{firstName: "venu", lastName: "surampudi"}}
	//proprietor.GetName()

	//	t.Fail()
}

/*
func TestEmployee_GetEmployeeFullName(t *testing.T) {
	fmt.Println("in test method q")
	type fields struct {
		firstName string
		lastName  string
		position  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"test1", fields{"surampudi", "venu", ""}, "surampudi,venu"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			employee := &Employee{
				firstName: tt.fields.firstName,
				lastName:  tt.fields.lastName,
				position:  tt.fields.position,
			}
			if got := employee.GetEmployeeFullName(); got != tt.want {
				t.Errorf("Employee.GetEmployeeFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}
*/
