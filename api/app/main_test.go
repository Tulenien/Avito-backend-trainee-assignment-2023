package main

import "testing"

func TestSample(t *testing.T) {
	if someFunc(2) != 4 {
		t.Error("2 + 2 != 4")
	}
}