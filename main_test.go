package main

import "testing"

func TestSum(t *testing.T) {
	a := 40
	b := 10

	total := sum(a, b)

	if(total != 50) {
		t.Errorf("Invalid result. Expected: %d, Got: %d", a + b, total)
	}
}