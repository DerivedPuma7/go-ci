package main

import "testing"

func TestSum(t *testing.T) {
	a := 20
	b := 10

	total := sum(a, b)

	if(total != 30) {
		t.Errorf("Invalid result. Expected: %d, Got: %d", a + b, total)
	}
}