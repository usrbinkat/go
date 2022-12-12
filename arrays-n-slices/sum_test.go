package main

import (
	"testing"
)

func TestSum(t *testing.T) {
	msg := "\nExpected: %d\nReturned: %d\nNumbers: %v"

	t.Run("collection of 5 numbers", func(t *testing.T) {
	    numbers := []int{1, 2, 3, 4, 5}
		want := 15

		got := Sum(numbers)
		if got!= want {
			t.Errorf(msg, got, want, numbers)
		}
    })

	t.Run("collection of 5 numbers", func(t *testing.T) {
	    numbers := []int{1, 2, 3}
		want := 6

		got := Sum(numbers)
		if got!= want {
			t.Errorf(msg, got, want, numbers)
		}
    })
}