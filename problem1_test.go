package main
import (
	"testing"
	"errors"
)

func TestProblem1(t *testing.T) {
	result := Problem1()
	if result != 233168 {
		t.Error(errors.New("Problem1 failed"))
	}
}