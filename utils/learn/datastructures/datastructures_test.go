package datastructures

import (
	"fmt"
	"testing"
)

func TestCreateAndIterateMap(t *testing.T) {
	m := CreateAndIterateMap()
	fmt.Println(m)
	if len(m) != 1 {
		t.Error("Failed")
	}
}

func TestLoopOverArray(t *testing.T) {
	fmt.Println("running ..")
	type args struct {
		inputArray []int
	}
	tests := []struct {
		name string
		args []int
	}{
		{"name", []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		fmt.Println("running ..")
		t.Run(tt.name, func(t *testing.T) {
			LoopOverArray(tt.args)
		})
	}
}
