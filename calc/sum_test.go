package calc_test

import (
	"study-ci/calc"
	"testing"
)

func TestSum(t *testing.T) {
	result := calc.Sum(10, 10)
	if result != 20 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}

func TestSub(t *testing.T) {
	result := calc.Sub(10, 10)
	if result != 0 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}

func TestMultiple(t *testing.T) {
	result := calc.Multiple(10, 10)
	if result != 100 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}

func TestDivise(t *testing.T) {
	result := calc.Divise(10, 10)
	if result != 1 {
		t.Errorf("Return is invalid. The correct value is %d", result)
	}
}
