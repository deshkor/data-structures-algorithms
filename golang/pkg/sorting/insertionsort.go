package sorting

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](vals []T) []T {
	for idx := 1; idx < len(vals); idx++ {
		key := vals[idx]

		j := idx - 1
		for ; j >= 0 && vals[j] > key; j-- {
			vals[j+1] = vals[j]
		}

		vals[j+1] = key
	}

	return vals
}