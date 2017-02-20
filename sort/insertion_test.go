package sort

import (
	"math/rand"
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
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
		if InsertionSort(tt.in); !reflect.DeepEqual(tt.in, tt.want) {
			t.Errorf("%v: InsertionSort(in) = %v, want %v", i, tt.in, tt.want)
		}
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	n := 1 << 10
	s := make([]byte, n)
	in := make([]int, n)
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		rand.Read(s)
		for j, b := range s {
			in[j] = int(b)
		}
		b.StartTimer()
		InsertionSort(in)
	}
}
