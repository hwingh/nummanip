package calc

import "testing"

func TestAdd(t *testing.T) {
	result, _ := Add(1, 2, 3)

	if result != 6 {
		t.Errorf("Add(1, 2, 3) FAILED, expected %v but got value %v", 6, result)
	} else {
		t.Logf("Add(1, 2, 3) PASSED, expected %v and got value %v", 6, result)
	}
}
