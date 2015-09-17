package main
import (
	"testing"
	"errors"
)

func TestProblem9(t *testing.T) {
	result := Problem9()
	if result != 31875000 {
		t.Error(errors.New("Problem9 failed"))
	}
}