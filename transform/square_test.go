package transform

import (
	"reflect"
	"testing"
)

func TestSquareSlices(t *testing.T) {
	testSlice := []int{1, 2, 3, 4, 5}
	expectResult := []int{1, 4, 9, 16, 25}

	result := SquareSlices(testSlice)

	if reflect.DeepEqual(expectResult, result) {
		t.Logf("SquareSlices PASSED")
	} else {
		t.Errorf("SquareSlices FAILED, expected %v but got %v", expectResult, result)
	}
}
