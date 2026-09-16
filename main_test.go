package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerator(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		size := 100
		slice := generateRandomElements(size)
		assert.NotNil(t, slice)
		assert.Len(t, slice, size)
	})
	t.Run("negative", func(t *testing.T) {
		slice := generateRandomElements(-1)
		require.Nil(t, slice)
	})
	t.Run("zero", func(t *testing.T) {
		slice := generateRandomElements(0)
		require.Nil(t, slice)
	})
}

func TestMaximumAndChunks(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{
			name:   "nil",
			input:  nil,
			output: 0,
		},
		{
			name:   "zero",
			input:  []int{},
			output: 0,
		}, {
			name:   "positive",
			input:  []int{24, 15, 42, 10, 7},
			output: 42,
		}, {
			name:   "negative",
			input:  []int{-50, -20, -3, -100, -24},
			output: -3,
		}, {
			name:   "single",
			input:  []int{42},
			output: 42,
		},
		{
			name:   "equal",
			input:  []int{42, 42, 42, 42, 42},
			output: 42,
		},
	}
	for _, el := range tests {
		t.Run(el.name, func(t *testing.T) {
			max := maximum(el.input)
			assert.Equal(t, el.output, max)
		})
	}
	t.Run("compare", func(t *testing.T) {
		slice := generateRandomElements(SIZE)
		maxBasic := maximum(slice)
		maxMultiFlow := maxChunks(slice)
		assert.Equal(t, maxBasic, maxMultiFlow)
	})
}
