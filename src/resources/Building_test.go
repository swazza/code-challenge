package resources

import (
	"math"
	"testing"
)

func TestAbs(t *testing.T) {
	got := math.Abs(-1)
	if got != 1.0 {
		//t.Errorf("Abs(-1) = %d; want 1", got)
	}
}