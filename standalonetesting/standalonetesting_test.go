package standalonetesting_test

import (
	"fmt"
	"testing"

	"github.com/BorisPlus/golang_notes/standalonetesting"
)

func TestPow2(t *testing.T) {
	x := 2
	y := standalonetesting.Pow2(x)
	if y != x*x {
		t.Errorf("Result was incorrect, got: %v, want: %v.", x*x, y)
	}
	fmt.Printf("Result is correct, got: %d. OK.\n", x*x)
}
