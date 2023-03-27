package insertionsort

import "golang.org/x/exp/constraints"

func InsertionSort[T constraints.Ordered](values []T) []T {
	for idx := 1; idx < len(values); idx++ {
		key := values[idx]

		j := idx - 1
		for ; j >= 0 && values[j] > key; j-- {
			values[j+1] = values[j]
		}

		values[j+1] = key
	}

	return values
}