package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertionSortInt(t *testing.T) {
	for _, tc := range testCasesInt {
		t.Run(tc.Name, func(t *testing.T){
			res := InsertionSort(tc.Values)
			assert.Equal(t, tc.Expected, res)
		})
	}
}

func TestInsertionSortFloat(t *testing.T) {
	for _, tc := range testCasesFloat32 {
		t.Run(tc.Name, func(t *testing.T){
			res := InsertionSort(tc.Values)
			assert.Equal(t, tc.Expected, res)
		})
	}
}