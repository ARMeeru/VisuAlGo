package sumnum_test

import (
	"math"
	"testing"

	sumnum "github.com/ARMeeru/VisuAlGo/problems/2.SumNum"
)

func TestSum(t *testing.T) {
	tests := []struct {
		name     string
		numbers  []int
		expected int
	}{
		{
			name:     "Sum of positive numbers",
			numbers:  []int{1, 2, 3, 4, 5},
			expected: 15,
		},
		{
			name:     "Sum of negative numbers",
			numbers:  []int{-1, -2, -3, -4, -5},
			expected: -15,
		},
		{
			name:     "Sum of mixed positive and negative numbers",
			numbers:  []int{-1, 2, -3, 4, -5},
			expected: -3,
		},
		{
			name:     "Sum of an empty list",
			numbers:  []int{},
			expected: 0,
		},
		{
			name:     "Sum of a single element list",
			numbers:  []int{7},
			expected: 7,
		},
		{
			name:     "Sum of a list with zero",
			numbers:  []int{0, 0, 0, 0},
			expected: 0,
		},
		{
			name:     "Sum of a list with large numbers",
			numbers:  []int{1000000, 2000000, 3000000},
			expected: 6000000,
		},
		{
			name:     "Sum of a list with maximum and minimum integer values",
			numbers:  []int{math.MaxInt64, math.MinInt64, 0},
			expected: -1,
		},
		{
			name: "Sum of a very large list",
			numbers: func() []int {
				largeList := make([]int, 1000000)
				for i := 0; i < 1000000; i++ {
					largeList[i] = 1
				}
				return largeList
			}(),
			expected: 1000000,
		},
		{
			name:     "Sum of two numbers that cancel each other out",
			numbers:  []int{math.MaxInt64, -math.MaxInt64},
			expected: 0,
		},
		{
			name:     "Alternating high positive and negative numbers",
			numbers:  []int{math.MaxInt64, -math.MaxInt64, math.MaxInt64, -math.MaxInt64},
			expected: 0,
		},
		{
			name:     "Nil input",
			numbers:  nil,
			expected: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sumnum.Sum(tt.numbers)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}
