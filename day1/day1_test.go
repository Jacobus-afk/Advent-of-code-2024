package main

import (
	"reflect"
	"testing"
)

func TestNumList(t *testing.T) {
	t.Run("test for smallest", func(t *testing.T) {
		numList := NumList{
			leftList:  []int{3, 4, 2, 1, 3, 9},
			rightList: []int{4, 3, 5, 3, 9, 3},
		}

		gotLeft, gotRight := numList.smallest()

		wantLeft := 1
		wantRight := 3

		if gotLeft != wantLeft {
			t.Errorf("got %d, want %d", gotLeft, wantLeft)
		}
		if gotRight != wantRight {
			t.Errorf("got %d, want %d", gotRight, wantRight)
		}
	})

	t.Run("test for distance between smallest", func(t *testing.T) {
		got := distance(1, 3)

		want := 2

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test for pop off array", func(t *testing.T) {
		numList := NumList{
			leftList:  []int{3, 4, 2, 1, 3, 9},
			rightList: []int{4, 3, 5, 3, 9, 3},
		}

		numList.popMin()

		want := NumList{
			leftList:  []int{3, 4, 2, 3, 9},
			rightList: []int{4, 5, 3, 9, 3},
		}

		if !reflect.DeepEqual(numList, want) {
			t.Errorf("got %d, want %d", numList, want)
		}
	})

	t.Run("test for sum of distances", func(t *testing.T) {
		num_list := NumList{
			leftList:  []int{3, 4, 2, 1, 3, 3},
			rightList: []int{4, 3, 5, 3, 9, 3},
		}

		got := num_list.DistanceSum()

		want := 11

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test for similarity", func(t *testing.T) {
    leftVal := 3
    rightList := []int{4, 3, 5, 3, 9, 3}

		got := getSimilarity(leftVal, rightList)

		want := 9

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("test for sum of similarities", func(t *testing.T) {
		num_list := NumList{
			leftList:  []int{3, 4, 2, 1, 3, 3},
			rightList: []int{4, 3, 5, 3, 9, 3},
		}

		got := num_list.SimilaritySum()

		want := 31

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
