package main

import (
	"errors"
	"testing"
)

func TestProblem12(t *testing.T) {
	result := Problem12()
	if result != 76576500 {
		t.Error(errors.New("Problem12 failed"))
	}
}
