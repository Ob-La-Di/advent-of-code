package main

import "testing"

func Test_splitFileContent(t *testing.T) {
	input := []byte{'h', 'e', 'l', 'l', 'o', '\n', 'w', 'o', 'r', 'l', 'd'}
	result := splitFileContent(input)

	if len(result) != 2 {
		t.Errorf("array length invalid: expected 2, got %d", len(result))
	}

	if result[0] != "hello" {
		t.Errorf("invalid item at index 0: expected 'hello', got %s", result[0])
	}

	if result[1] != "world" {
		t.Errorf("invalid item at index 1: expected 2, got %s", result[1])
	}
}

func Test_convertArrayToInt(t *testing.T) {
	test := []struct {
		input    []string
		expected []int
	}{{
		input:    []string{"123", "456", "789"},
		expected: []int{123, 456, 789},
	},
	}

	for _, item := range test {
		result, err := convertArrayToInt(item.input)

		if err != nil {
			t.Errorf("expected nil error, got: %v", err)
		}

		if len(result) != len(item.input) {
			t.Errorf("array length missmatch, expected %d, got %d", len(item.input), len(result))
		}

		for index, number := range result {
			if item.expected[index] != number {
				t.Errorf("wrong converted value at index %d, expected %d, got %d", index, item.expected[index], number)
			}
		}
	}

	result, err := convertArrayToInt([]string{"789", "hello", "654"})
	if err == nil {
		t.Errorf("expected error to be not nil, got result with value: %v", result)
	}
}

func Test_getNumberMeasurements(t *testing.T) {
	params := []struct {
		input    []int
		expected int
	}{
		{
			input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 8,
		},
		{
			input:    []int{9, 8, 7, 6, 5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			input:    []int{1, 2, 3, 6, 5, 4, 9, 8, 7},
			expected: 4,
		},
	}

	for _, item := range params {
		result := getNumberOfMeasurementsGreaterThanPrevious(item.input)
		if result != item.expected {
			t.Errorf("invalid result value: expected: %d, got %d", result, item.expected)
		}
	}
}

func Test_groupByMeasurementWindow(t *testing.T) {
	params := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			expected: []int{607, 618, 618, 617, 647, 716, 769, 792},
		},
		{
			input:    []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263, 215, 78},
			expected: []int{607, 618, 618, 617, 647, 716, 769, 792, 738, 556},
		},
	}

	for _, item := range params {
		result := groupByMeasurementWindow(item.input, 3)
		if len(result) != len(item.expected) {
			t.Errorf("invalid result array length: expected: %d, got %d", len(item.expected), len(result))
		}

		for index, number := range result {
			if item.expected[index] != number {
				t.Errorf("wrong converted value at index %d, expected %d, got %d", index, item.expected[index], number)
			}
		}
	}
}
