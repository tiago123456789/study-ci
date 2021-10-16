package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	result := sum(10, 10)
	if result != 20 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}
