package main

import "testing"

func TestSum(t *testing.T) {
	a := 10
	b := 20

	total := sum(a, b)

	if(total != 30) {
		t.Errorf("Invalid result. Expected: %d, Got: %d", a + b, total)
	}
}