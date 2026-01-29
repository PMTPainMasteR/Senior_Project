// simulation_test.go
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
