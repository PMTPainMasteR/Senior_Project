package main

import (
	"testing"
)

func TestSimulationRuns(t *testing.T) {
	n1, trans := RunSimulationForTest(100)

	if len(n1) == 0 {
		t.Error("n1 map is empty")
	}
	total := trans["P10"] + trans["P01"] + trans["P1T"] + trans["P0T"]
	if total == 0 {
		t.Error("No transitions recorded")
	}
}

func TestDisplayProbabilities(t *testing.T) {
	n1 := map[int]int{0: 50, 1: 50}
	total := 100
	displayProbabilities(n1, total)
}

func TestSecureFloat64(t *testing.T) {
	val, err := secureFloat64()
	if err != nil {
		t.Errorf("secureFloat64 returned error: %v", err)
	}
	if val < 0.0 || val >= 1.0 {
		t.Errorf("secureFloat64 returned out of range [0,1): %v", val)
	}
}

func TestMainExecution(t *testing.T) {
	main()
}
