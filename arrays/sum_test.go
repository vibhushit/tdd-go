package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	expected := []int{3, 9}

	if !slices.Equal(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("If array is not empty", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		expected := []int{2, 9}

		//t.Log("got", got, "expected", expected)
		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)

		}
	})

	t.Run("If arr is empty", func(t *testing.T) {
		got := SumAllTails([]int{})
		expected := []int{0}

		if !slices.Equal(got, expected) {
			t.Errorf("got %v expected %v", got, expected)
		}
	})
}
