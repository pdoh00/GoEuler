package main
import (
	"testing"
	"errors"
)

func TestProblem2(t *testing.T) {
	result := Problem2()
	if result != 4613732 {
		t.Error(errors.New("Problem2 failed"))
	}
}