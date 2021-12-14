package main

import "testing"

func Test_readLine(t *testing.T) {
	params := []struct {
		input    string
		expected movement
	}{{
		input: "down 1",
		expected: movement{
			down,
			1,
		},
	}, {
		input: "up 1",
		expected: movement{
			up,
			1,
		},
	}, {
		input: "forward 1",
		expected: movement{
			forward,
			1,
		},
	}}

	for _, item := range params {
		result, err := readLine(item.input)

		if err != nil {
			t.Errorf("expected err to be nil, got %v", err)
		}
		if item.expected.dir != result.dir {
			t.Errorf("wrong movement dir, expected %s, got %s", item.expected.dir, result.dir)
		}
		if item.expected.value != result.value {
			t.Errorf("wrong movement value, expected %d, got %d", item.expected.value, result.value)
		}
	}

	paramsError := []struct {
		input    string
		expected string
	}{{
		input:    "down abc",
		expected: "invalid number in value: abc",
	}, {
		input:    "up def",
		expected: "invalid number in value: def",
	}, {
		input:    "forward ghi",
		expected: "invalid number in value: ghi",
	}, {
		input:    "totally irrelevant input",
		expected: "expected only one space in input string",
	}, {
		input:    "whatever 78",
		expected: "invalid direction: whatever",
	},
	}

	for _, item := range paramsError {
		result, err := readLine(item.input)

		if err == nil {
			t.Errorf("expected err not to be nil, got result with value: %v", result)
		}
		if err.Error() != item.expected {
			t.Errorf("wrong error value, got %s", err.Error())
		}
	}

}
