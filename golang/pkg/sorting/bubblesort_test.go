package sorting

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSortInt(t *testing.T) {
	for _, tc := range testCasesInt {
		t.Run(tc.Name, func(t *testing.T){
			res := BubbleSort(tc.Values)
			assert.Equal(t, tc.Expected, res)
		})
	}
}

func TestBubbleSortFloat(t *testing.T) {
	for _, tc := range testCasesFloat32 {
		t.Run(tc.Name, func(t *testing.T){
			res := BubbleSort(tc.Values)
			assert.Equal(t, tc.Expected, res)
		})
	}
}