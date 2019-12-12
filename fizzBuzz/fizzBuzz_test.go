package fizzBuzz

import (
	"fmt"
	"testing"
)

func Test_SayFizzBuzz_Input_1_Should_Be_1(t *testing.T) {
	expected := "1"
	actual := sayFizzBuzz(1)
	if expected != actual {
		fmt.Printf("Expected %s but got %s", expected, actual)
	}
}
