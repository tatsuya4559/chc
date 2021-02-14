package main

import (
	"reflect"
	"testing"
)

func TestDetermineCase(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "foo", expected: SNAKE},
		{input: "FOO", expected: UPPER},
		{input: "foo_bar", expected: SNAKE},
		{input: "FOO_BAR", expected: UPPER},
		{input: "fooBar", expected: CAMEL},
		{input: "FooBar", expected: PASCAL},
	}

	for _, tt := range tests {
		if determineCase(tt.input) != tt.expected {
			t.Errorf("determineCase(%q) was %q, want %q",
				tt.input, determineCase(tt.input), tt.expected)
		}
	}
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{input: "foo", expected: []string{"foo"}},
		{input: "FOO", expected: []string{"foo"}},
		{input: "foo_bar", expected: []string{"foo", "bar"}},
		{input: "FOO_BAR", expected: []string{"foo", "bar"}},
		{input: "fooBar", expected: []string{"foo", "bar"}},
		{input: "FooBar", expected: []string{"foo", "bar"}},
	}

	for _, tt := range tests {
		actual := tokenize(tt.input)
		if !reflect.DeepEqual(actual, tt.expected) {
			t.Errorf("tokenize(%q) was %v, want %v",
				tt.input, actual, tt.expected)
		}
	}
}

func TestToSnake(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "foo", expected: "foo"},
		{input: "foo_bar", expected: "foo_bar"},
		{input: "FOO_BAR", expected: "foo_bar"},
		{input: "fooBar", expected: "foo_bar"},
		{input: "FooBar", expected: "foo_bar"},
	}

	for _, tt := range tests {
		actual := toSnake(tt.input)
		if actual != tt.expected {
			t.Errorf("toSnake(%q) was %q, want %q",
				tt.input, actual, tt.expected)
		}
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "foo", expected: "FOO"},
		{input: "foo_bar", expected: "FOO_BAR"},
		{input: "FOO_BAR", expected: "FOO_BAR"},
		{input: "fooBar", expected: "FOO_BAR"},
		{input: "FooBar", expected: "FOO_BAR"},
	}

	for _, tt := range tests {
		actual := toUpper(tt.input)
		if actual != tt.expected {
			t.Errorf("toUpper(%q) was %q, want %q",
				tt.input, actual, tt.expected)
		}
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "foo", expected: "foo"},
		{input: "foo_bar", expected: "fooBar"},
		{input: "FOO_BAR", expected: "fooBar"},
		{input: "fooBar", expected: "fooBar"},
		{input: "FooBar", expected: "fooBar"},
	}

	for _, tt := range tests {
		actual := toCamel(tt.input)
		if actual != tt.expected {
			t.Errorf("toCamel(%q) was %q, want %q",
				tt.input, actual, tt.expected)
		}
	}
}

func TestToPascal(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "foo", expected: "Foo"},
		{input: "foo_bar", expected: "FooBar"},
		{input: "FOO_BAR", expected: "FooBar"},
		{input: "fooBar", expected: "FooBar"},
		{input: "FooBar", expected: "FooBar"},
	}

	for _, tt := range tests {
		actual := toPascal(tt.input)
		if actual != tt.expected {
			t.Errorf("toPascal(%q) was %q, want %q",
				tt.input, actual, tt.expected)
		}
	}
}
