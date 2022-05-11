package homework

import (
	"reflect"
	"testing"
)

func TestSortMapValues(t *testing.T) {
	input := map[int]string{2: "a", 0: "b", 1: "c"}
	got := sortMapValues(input)
	want := []string{"b", "c", "a"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}

	input2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	got = sortMapValues(input2)
	want = []string{"bb", "aa", "cc"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
