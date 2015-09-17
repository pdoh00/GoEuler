package main
import (
	"testing"
	"errors"
)

func TestProblem3(t *testing.T) {
	result := Problem3()
	if result != 6857 {
		t.Error(errors.New("Problem3 failed"))
	}
}