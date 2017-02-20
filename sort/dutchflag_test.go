package sort

import (
	"reflect"
	"testing"
)

var params = []struct {
	s    []int
	a, b int
	want []int
}{
	{[]int{2, 7, 1, -92}, 0, 3, []int{-92, 2, 1, 7}},
}

func TestThreeWayPartition(t *testing.T) {
	for _, p := range params {
		test(t, p.s, p.a, p.b, p.want)
	}
}

func test(t *testing.T, s []int, a, b int, want []int) {
	if ThreeWayPartition(s, a, b); !reflect.DeepEqual(s, want) {
		t.Errorf("ThreeWayPartition(s, %v, %v) = %v, want %v", a, b, s, want)
	}
}
