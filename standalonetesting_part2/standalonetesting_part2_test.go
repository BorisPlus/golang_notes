package standalonetesting_part2_test

import (
	"fmt"
	"testing"

	"github.com/BorisPlus/golang_notes/standalonetesting_part2"
)

func TestPow2(t *testing.T) {
	x := 2
	y := standalonetesting_part2.Pow3(x)
	if y != x*x*x {
		t.Errorf("Result was incorrect, got: %v, want: %v.", x*x*x, y)
	}
	fmt.Printf("Result is correct, got: %d^3 = %d. OK.\n", x , x*x*x)
}
