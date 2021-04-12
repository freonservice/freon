package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqueStringSlice(t *testing.T) {
	var params = []struct {
		value    []string
		expected []string
	}{
		{
			value:    []string{"1", "1", "1"},
			expected: []string{"1"},
		},
		{
			value:    []string{"A", "B", "C"},
			expected: []string{"A", "B", "C"},
		},
		{
			value:    []string{"A", ""},
			expected: []string{"A"},
		},
	}

	for _, p := range params {
		res := uniqueStringSlice(p.value)
		assert.Equal(t, p.expected, res)
	}
}

func TestCreateConcatenatedString(t *testing.T) {
	var params = []struct {
		value    []string
		expected string
	}{
		{
			value:    []string{"1", "1", "1"},
			expected: "1",
		},
		{
			value:    []string{"A", "B", "C"},
			expected: "A,B,C",
		},
		{
			value:    []string{"C", "B", "A"},
			expected: "A,B,C",
		},
		{
			value:    []string{"service", "translations", "affordable"},
			expected: "affordable,service,translations",
		},
		{
			value:    []string{"service", "translations", "affordable", "affordable"},
			expected: "affordable,service,translations",
		},
	}

	for _, p := range params {
		res := createConcatenatedString(p.value)
		assert.Equal(t, p.expected, res)
	}
}
