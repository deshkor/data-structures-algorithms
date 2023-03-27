package sorting

import "golang.org/x/exp/constraints"

func BubbleSort[T constraints.Ordered](vals []T) []T {
	for {
		swapped := false

		for i := 1; i < len(vals); i++ {
			cur := vals[i]
			if vals[i-1] > vals[i] {
				vals[i] = vals[i-1]
				vals[i-1] = cur

				swapped = true
			}
		}

		if !swapped {
			break
		}
	}

	return vals
}