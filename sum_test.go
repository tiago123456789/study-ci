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

func TestSub(t *testing.T) {
	result := sub(10, 10)
	if result != 0 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}

func TestMultiple(t *testing.T) {
	result := multiple(10, 10)
	if result != 100 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}

func TestDivise(t *testing.T) {
	result := divise(10, 10)
	if result != 1 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}
