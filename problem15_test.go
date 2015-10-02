package main

import (
	"testing"
	"errors"
	"fmt"
)

func TestProblem15(t *testing.T){
	result := Problem15()

	if result != int64(137846528820) {
		t.Error(errors.New(fmt.Sprintf("Problem15 failed. Expected 137846528820 Received %d", result)))
	}
}