package main

import (
	"errors"
	"testing"
	"fmt"
)

func TestProblem14(t *testing.T) {
	result := Problem14()

	if result != 837799 {
		t.Error(errors.New(fmt.Sprintf("Problem13 failed. Expected 837799 Received %d", result)))
	}
}

