package main

import (
	"reflect"
	"testing"
)

// TestCleanJSON tests the cleanJSON function with various inputs
func TestCleanJSON(t *testing.T) {
	testCases := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name:     "remove empty strings and nil from map",
			input:    map[string]any{"a": "", "b": nil, "c": "value"},
			expected: map[string]any{"c": "value"},
		},
		{
			name:     "remove empty strings and nil from slice",
			input:    []any{"", nil, "value", 123},
			expected: []any{"value", 123},
		},
		{
			name:     "nested structures",
			input:    map[string]any{"a": map[string]any{"b": "", "c": "nested"}, "d": ""},
			expected: map[string]any{"a": map[string]any{"c": "nested"}},
		},
		{
			name:     "entire map becomes empty",
			input:    map[string]any{"a": map[string]any{"b": ""}, "c": ""},
			expected: nil,
		},
		{
			name:     "entire slice becomes empty",
			input:    []any{"", nil},
			expected: nil,
		},
		{
			name:     "non-empty string",
			input:    "non-empty",
			expected: "non-empty",
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
		{
			name:     "number",
			input:    123,
			expected: 123,
		},
		{
			name:     "nil",
			input:    nil,
			expected: nil,
		},
		{
			name:     "remove `refresh: false` from root",
			input:    map[string]any{"refresh": false, "a": "foo"},
			expected: map[string]any{"a": "foo"},
		},
		{
			name:     "don't remove non-root `refresh: false`",
			input:    map[string]any{"a": map[string]any{"refresh": false, "a": "bar"}, "b": "foo"},
			expected: map[string]any{"a": map[string]any{"refresh": false, "a": "bar"}, "b": "foo"},
		},
		{
			name:     "don't remove refresh string",
			input:    map[string]any{"refresh": "1m", "a": "foo"},
			expected: map[string]any{"refresh": "1m", "a": "foo"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := cleanJSON(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, got %v", tc.expected, result)
			}
		})
	}
}
