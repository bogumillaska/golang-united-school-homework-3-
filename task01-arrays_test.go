package homework

import "testing"

func TestAverage(t *testing.T) {
	got := average([15]float32{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	want := 1.4

	if got != float32(want) {
		t.Errorf("got %v, wanted %v", got, want)
	}
	return
}
