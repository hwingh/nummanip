package calc

import "testing"

type TestDataItem struct {
	inputs   []int
	result   int
	hasError bool
}

func TestAdd(t *testing.T) {

	dataItems := []TestDataItem{
		{[]int{1, 2, 3}, 6, false},
		{[]int{99, 99}, 198, false},
		{[]int{1, 1, 5, 5, 6}, 18, false},
		{[]int{1}, 0, false},
	}

	for _, item := range dataItems {
		result, err := Add(item.inputs...)

		if item.hasError {
			if err == nil {
				t.Errorf("Add() with args %v: FAILED, expected an error but got value %v", item.inputs, result)
			} else {
				t.Logf("Add() with args %v: PASSED, expected an error and got an error %v", item.inputs, err)
			}
		} else {
			if result != item.result {
				t.Errorf("Add() with args %v: FAILED, expected %v but got value %v", item.inputs, item.result, result)
			} else {
				t.Logf("Add() with args %v: PASSED, expected %v and got %v", item.inputs, item.result, result)
			}
		}
	}

}
