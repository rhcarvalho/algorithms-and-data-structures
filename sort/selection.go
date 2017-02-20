package sort

func SelectionSort(s []int) {
	// i: current item, all items up to s[i-1] are sorted.
	for i := 0; i < len(s); i++ {
		// m: index of the smallest unsorted item in s.
		m := i
		// recompute m
		for j := i; j < len(s); j++ {
			if s[j] < s[m] {
				m = j
			}
		}
		s[i], s[m] = s[m], s[i]
	}
}
