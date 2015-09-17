package main
import (
	"testing"
	"errors"
)

func TestProblem10(t *testing.T) {
	result := Problem10()
	if result != 142913828922 {
		t.Error(errors.New("Problem10 failed"))
	}
}