package main
import (
	"testing"
	"errors"
)

func TestProblem11(t *testing.T) {
	result := Problem11()
	if result != 70600674 {
		t.Error(errors.New("Problem11 failed"))
	}
}