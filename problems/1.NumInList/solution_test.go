package numinlist_test

import (
	"testing"

	numinlist "github.com/ARMeeru/VisuAlGo/problems/1.NumInList"
)

func TestNumInList(t *testing.T) {
	tests := []struct {
		name     string
		list     []int
		num      int
		expected bool
	}{
		{
			name:     "Number is present in the list",
			list:     []int{1, 2, 3, 4, 5},
			num:      3,
			expected: true,
		},
		{
			name:     "Number is not present in the list",
			list:     []int{1, 2, 3, 4, 5},
			num:      6,
			expected: false,
		},
		{
			name:     "Empty list",
			list:     []int{},
			num:      1,
			expected: false,
		},
		{
			name:     "Single element list, number present",
			list:     []int{7},
			num:      7,
			expected: true,
		},
		{
			name:     "Single element list, number not present",
			list:     []int{7},
			num:      3,
			expected: false,
		},
		{
			name:     "Large list, number present",
			list:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			num:      10,
			expected: true,
		},
		{
			name:     "Negative numbers in the list, number present",
			list:     []int{-10, -5, -3, -1, 0, 2, 3, 7},
			num:      -5,
			expected: true,
		},
		{
			name:     "Negative numbers in the list, number not present",
			list:     []int{-10, -5, -3, -1, 0, 2, 3, 7},
			num:      8,
			expected: false,
		},
		{
			name:     "Mixed positive and negative numbers, number present",
			list:     []int{-100, -50, 0, 50, 100},
			num:      0,
			expected: true,
		},
		{
			name:     "Duplicate numbers in the list, number present",
			list:     []int{1, 2, 3, 3, 4, 5},
			num:      3,
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := numinlist.NumInList(tt.list, tt.num)
			if result != tt.expected {
				t.Errorf("got %v, want %v", result, tt.expected)
			}
		})
	}
}