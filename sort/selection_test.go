package sort

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	tests := []struct {
		in   []int
		want []int
	}{
		{
			[]int{5, 4, 3, 7, 1, -1},
			[]int{-1, 1, 3, 4, 5, 7},
		},
	}
	for i, tt := range tests {
		if SelectionSort(tt.in); !reflect.DeepEqual(tt.in, tt.want) {
			t.Errorf("%v: SelectionSort(in) = %v, want %v", i, tt.in, tt.want)
		}
	}
}
