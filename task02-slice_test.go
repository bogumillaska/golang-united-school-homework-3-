package homework

import (
	"testing"
)

func TestReverse(t *testing.T) {
	input := []int64{1, 2, 5, 15}
	got := reverse(input)
	want := []int64{15, 5, 2, 1}

	if len(got) != len(want) {
		t.Errorf("got %v, wanted %v", got, want)
	}

	for i, j := 0, len(input)-1; j >= 0; i, j = i+1, j-1 {
		if want[i] != input[j] {
			t.Errorf("wrong")
		}
	}
	return
}
