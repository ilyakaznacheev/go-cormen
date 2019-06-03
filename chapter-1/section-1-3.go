package chapter

import "math"

// Section 1.3 Designing algorithms

const maxInt = int(1 >> 0)

// Merge the two sorted subsequences to produce the sorted answer
func Merge(A []int, p, q, r int) {
	n1 := q - p + 1
	n2 := r - q
	L := make([]int, n1, n1)
	R := make([]int, n2, n2)

	var i, j int

	for i = 0; i < n1-1; i++ {
		L[i] = A[p+i-1]
	}

	for j = 0; j < n2-1; j++ {
		R[j] = A[q+j]
	}
	L[n1] = maxInt
	R[n2] = maxInt

	i, j = 0, 0

	for k := p; k < r; k++ {
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}
}

// MergeSort sorts the elements in the subarray
func MergeSort(A []int, p, r int) {
	if p < r {
		q := int(math.Floor(float64(p+r) / 2.0))
		MergeSort(A, p, q)
		MergeSort(A, q+1, r)
		Merge(A, p, q, r)
	}
}
