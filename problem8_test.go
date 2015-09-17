package main
import (
	"testing"
	"errors"
)

func TestProblem8(t *testing.T) {
	result := Problem8()
	if result != 23514624000 {
		t.Error(errors.New("Problem8 failed"))
	}
}