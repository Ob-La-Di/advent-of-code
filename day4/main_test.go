package main

import "testing"

func Test_parseRandomNumbers(t *testing.T) {
	params := []struct {
		input    []byte
		expected []int
	}{
		{
			input:    []byte{'1', ',', '2', ',', '3'},
			expected: []int{1, 2, 3},
		},
		{
			input:    []byte{'1'},
			expected: []int{1},
		},
		{
			input:    []byte{},
			expected: []int{},
		},
	}

	for _, item := range params {
		result, err := parseRandomNumbers(item.input)
		if err != nil {
			t.Errorf("expected err to be nil, got %v", err)
		}
		if len(item.expected) != len(result) {
			t.Errorf("arrays size differ, expected %d got %d", len(item.expected), len(result))
		}

		for i := 0; i < len(item.expected); i++ {
			if item.expected[i] != result[i] {
				t.Errorf("expected item at index %d to be %d, got %d", i, item.expected[i], result[i])
			}
		}
	}
}
