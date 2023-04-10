package sorting

import (
	"golang.org/x/exp/constraints"
)

func MergeSort[T constraints.Ordered](vals []T) []T {
	if len(vals) < 2 {
		return vals
	}

	mid := len(vals)/2
	valsLeft := MergeSort(vals[:mid])
	valsRight := MergeSort(vals[mid:])
	return mergeSort(valsLeft, valsRight)
}

func mergeSort[T constraints.Ordered](valsLeft, valsRight []T) []T {
	res := []T{}

	i := 0
	for i < len(valsLeft) {
		valLeft := valsLeft[i]
		reset := false

		for idx, valRight := range valsRight {
			if valRight < valLeft {
				res = append(res, valRight)

				if len(valsRight) == 1 {
					valsRight = []T{}
				} else {
					valsRight = append(valsRight[:idx], valsRight[idx+1:]...)
				}

				reset = true
				break
			}
		}

		if reset {
			continue
		}

		res = append(res, valLeft)
		i++
	}

	for _, valRight := range valsRight {
		res = append(res, valRight)
	}

	return res
}