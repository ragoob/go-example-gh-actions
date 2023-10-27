package main

import "testing"

func TestSum(t *testing.T) {
	expected := 10
	a := 5
	b := 5
	res := Sum(a, b)

	if expected != res {
		t.Errorf("expected %d but got %d", expected, res)
	}
}
