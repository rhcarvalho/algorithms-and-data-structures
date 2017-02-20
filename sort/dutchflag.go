package sort

import "log"

// ThreeWayPartition rearranges s such that elements in the range (-inf, a) come
// first, [a, b) next, and [b, +inf) last, preserving the order of elements in
// the same partition.
// Though not the same, this is based on what is described in the Dutch national
// flag problem: https://en.wikipedia.org/wiki/Dutch_national_flag_problem
func ThreeWayPartition(s []int, a, b int) {
	// lo: boundary index of ints < a
	// i: index of the current int under consideration
	// up: boundary index of ints > b
	// INVARIANT: lo <= i <= up
	lo, i, up := 0, 0, len(s)-1
	log.Printf("input:\n\ts: %v\n\ta: %v\n\tb: %v", s, a, b)
	for i <= up {
		switch {
		case s[i] < a:
			log.Printf("visiting %v:\tmove left", s[i])
			s[lo], s[i] = s[i], s[lo]
			lo++
			i++
		case s[i] > b:
			log.Printf("visiting %v:\tmove right", s[i])
			s[up], s[i] = s[i], s[up]
			up--
		default:
			log.Printf("visiting %v:\tgo to next", s[i])
			i++
		}
	}
	// reverse upper group to bring back the original order
	for i, j := up+1, len(s)-1; i < j; {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
