package main

import "testing"

func TestSum(t *testing.T) {
	expected := 87
	a := 33
	b := 54
	res := Sum(a, b)

	if expected != res {
		t.Errorf("expected %d but got %d", expected, res)
	}
}
