package chapter

// Section 1.1 Algorithms

// InsertionSort takes as a parameter an array A[0..n] containing a sequence of length n that is to be sorted. (In the code, the number n of elements in A is denoted by length[A].) The input numbers are sorted in place: the numbers are rearranged within the array A, with at most a constant number of them stored outside the array at any time. The input array A contains the sorted output sequence when InsertionSort is finished.
func InsertionSort(a []int) {
	var i, j int
	for j = 1; j < len(a); j++ {
		key := a[j]
		// Insert A[j] into the sorted sequence A[0 . . j - 1].
		i = j - 1
		for i >= 0 && a[i] > key {
			a[i+1] = a[i]
			i--
		}
		a[i+1] = key
	}
}
