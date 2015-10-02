package main

import(
	"testing"
	"errors"
	"fmt"
)

func TestProblem16(t *testing.T){
	result:=Problem16()
	if result != int64(1366) {
		t.Error(errors.New(fmt.Sprintf("Problem16 failed. Expected %d Received %d", 1366, result)))
	}
}