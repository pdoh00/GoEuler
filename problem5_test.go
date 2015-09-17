package main
import (
	"testing"
	"errors"
)

func TestProblem5(t *testing.T) {
	result := Problem5()
	if result != 232792560 {
		t.Error(errors.New("Problem5 failed"))
	}
}