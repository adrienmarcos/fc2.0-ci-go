package main

import "testing"

func TestSum(t *testing.T) {
	result := Add(15, 15)
	if result != 30 {
		t.Errorf("Expected %d, got %d", 30, result)
	}
}
