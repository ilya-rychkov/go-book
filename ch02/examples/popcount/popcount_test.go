package popcount

import "testing"

func TestPopCount(t *testing.T) {
	if PopCount(10) != 2 {
		t.Errorf("PopCount(10) = %d, want 2", PopCount(10))
	}
}
