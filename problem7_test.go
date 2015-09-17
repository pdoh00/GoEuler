package main
import (
	"testing"
	"errors"
)

func TestProblem7(t *testing.T) {
	result := Problem7()
	if result != 104743 {
		t.Error(errors.New("Problem7 failed"))
	}
}