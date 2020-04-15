package inheritance

import (
	"fmt"
	"testing"
)

func TestShapeOperation(t *testing.T) {
	fmt.Println("logged")
	t.Log("logged")
	var sop ShapeOperatons
	var r Rectange
	sop = r
	fmt.Printf(" sop type is %T and valud is %v", sop, sop)
}
