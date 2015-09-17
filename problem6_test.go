package main
import (
	"testing"
	"errors"
)

func TestProblem6(t *testing.T) {
	result := Problem6()
	if result != 25164150 {
		t.Error(errors.New("Problem6 failed"))
	}
}