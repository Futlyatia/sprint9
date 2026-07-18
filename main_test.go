package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateRandomElements(t *testing.T) {
	// prepare
	sizes := []int{0, 1, 10, 100}

	for _, size := range sizes {
		result := generateRandomElements(size)

		require.Len(t, result, size)
	}
}
func TestGenerateRandomElementsNegative(t *testing.T) {
	result := generateRandomElements(-1)

	require.Empty(t, result)
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{
			name: "empty slice",
			data: []int{},
			want: 0,
		},
		{
			name: "one element",
			data: []int{10},
			want: 10,
		},
		{
			name: "positive numbers",
			data: []int{1, 5, 3, 8, 2},
			want: 8,
		},
		{
			name: "negative numbers",
			data: []int{-1, -5, -3, -8, -2},
			want: -1,
		},
		{
			name: "already sorted",
			data: []int{1, 2, 3, 4, 5},
			want: 5,
		},
		{
			name: "reverse sorted",
			data: []int{5, 4, 3, 2, 1},
			want: 5,
		},
		{
			name: "equal values",
			data: []int{7, 7, 7, 7},
			want: 7,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// action
			result := maximum(test.data)

			// check
			assert.Equal(t, test.want, result)
		})
	}
}
