package main
import (
	"testing"
	"errors"
)

func TestProblem4(t *testing.T) {
	result := Problem4()
	if result != 906609 {
		t.Error(errors.New("Problem4 failed"))
	}
}