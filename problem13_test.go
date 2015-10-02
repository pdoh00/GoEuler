package main

import (
	"errors"
	"testing"
	"fmt"
)

func TestProblem13(t *testing.T) {
	result := Problem13()

	if result != int64(5537376230) {
		t.Error(errors.New(fmt.Sprintf("Problem13 failed. Expected 5537376230 Received %d", result)))
	}
}
