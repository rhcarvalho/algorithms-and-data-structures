package sort

func InsertionSort(s []int) {
	// i: current item, all items up to s[i-1] are sorted.
	for i := 1; i < len(s); i++ {
		// swap elements
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}
