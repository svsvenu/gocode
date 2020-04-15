package variables

import (
	"fmt"
	"testing"
)

func TestVariables(test *testing.T) {
	a, b := returnMultipleValues()
	fmt.Printf("a is %d", a)
	fmt.Println(b)
}
