package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeSortInt(t *testing.T) {
	for _, tc := range testCasesInt {
		t.Run(tc.Name, func(t *testing.T){
			res := MergeSort(tc.Values)
			assert.Equal(t, tc.Expected, res)
		})
	}
}

func TestMergeSortFloat(t *testing.T) {
	for _, tc := range testCasesFloat32 {
		t.Run(tc.Name, func(t *testing.T){
			res := MergeSort(tc.Values)
			assert.Equal(t, tc.Expected, res)
		})
	}
}