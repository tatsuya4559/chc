package main

import (
	"bytes"
	"reflect"
	"strings"
	"testing"
)

func TestDetermineCase(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "foo", want: SNAKE},
		{input: "FOO", want: UPPER},
		{input: "foo_bar", want: SNAKE},
		{input: "FOO_BAR", want: UPPER},
		{input: "fooBar", want: CAMEL},
		{input: "FooBar", want: PASCAL},
		{input: "foo-bar", want: LISP},
	}

	for _, tt := range tests {
		if determineCase(tt.input) != tt.want {
			t.Errorf("determineCase(%q) was %q, want %q",
				tt.input, determineCase(tt.input), tt.want)
		}
	}
}

func TestTokenize(t *testing.T) {
	tests := []struct {
		input string
		want  []string
	}{
		{input: "foo", want: []string{"foo"}},
		{input: "FOO", want: []string{"foo"}},
		{input: "foo_bar", want: []string{"foo", "bar"}},
		{input: "FOO_BAR", want: []string{"foo", "bar"}},
		{input: "fooBar", want: []string{"foo", "bar"}},
		{input: "FooBar", want: []string{"foo", "bar"}},
		{input: "foo-bar", want: []string{"foo", "bar"}},
	}

	for _, tt := range tests {
		got := tokenize(tt.input)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("tokenize(%q) was %v, want %v",
				tt.input, got, tt.want)
		}
	}
}

func TestToSnake(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "foo", want: "foo"},
		{input: "foo_bar", want: "foo_bar"},
		{input: "FOO_BAR", want: "foo_bar"},
		{input: "fooBar", want: "foo_bar"},
		{input: "FooBar", want: "foo_bar"},
		{input: "foo-bar", want: "foo_bar"},
	}

	for _, tt := range tests {
		got := toSnake(tt.input)
		if got != tt.want {
			t.Errorf("toSnake(%q) was %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestToUpper(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "foo", want: "FOO"},
		{input: "foo_bar", want: "FOO_BAR"},
		{input: "FOO_BAR", want: "FOO_BAR"},
		{input: "fooBar", want: "FOO_BAR"},
		{input: "FooBar", want: "FOO_BAR"},
		{input: "foo-bar", want: "FOO_BAR"},
	}

	for _, tt := range tests {
		got := toUpper(tt.input)
		if got != tt.want {
			t.Errorf("toUpper(%q) was %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestToCamel(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "foo", want: "foo"},
		{input: "foo_bar", want: "fooBar"},
		{input: "FOO_BAR", want: "fooBar"},
		{input: "fooBar", want: "fooBar"},
		{input: "FooBar", want: "fooBar"},
		{input: "foo-bar", want: "fooBar"},
	}

	for _, tt := range tests {
		got := toCamel(tt.input)
		if got != tt.want {
			t.Errorf("toCamel(%q) was %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestToPascal(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "foo", want: "Foo"},
		{input: "foo_bar", want: "FooBar"},
		{input: "FOO_BAR", want: "FooBar"},
		{input: "fooBar", want: "FooBar"},
		{input: "FooBar", want: "FooBar"},
		{input: "foo-bar", want: "FooBar"},
	}

	for _, tt := range tests {
		got := toPascal(tt.input)
		if got != tt.want {
			t.Errorf("toPascal(%q) was %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestToLisp(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "foo", want: "foo"},
		{input: "foo_bar", want: "foo-bar"},
		{input: "FOO_BAR", want: "foo-bar"},
		{input: "fooBar", want: "foo-bar"},
		{input: "FooBar", want: "foo-bar"},
		{input: "foo-bar", want: "foo-bar"},
	}

	for _, tt := range tests {
		got := toLisp(tt.input)
		if got != tt.want {
			t.Errorf("toLisp(%q) was %q, want %q",
				tt.input, got, tt.want)
		}
	}
}

func TestChangeCase(t *testing.T) {
	tests := []struct {
		opt  string
		word string
		want string
	}{
		{opt: "s", word: "FooBar", want: "foo_bar"},
		{opt: "snake", word: "FooBar", want: "foo_bar"},
		{opt: "u", word: "foo_bar", want: "FOO_BAR"},
		{opt: "upper", word: "foo_bar", want: "FOO_BAR"},
		{opt: "c", word: "foo_bar", want: "fooBar"},
		{opt: "camel", word: "foo_bar", want: "fooBar"},
		{opt: "p", word: "foo_bar", want: "FooBar"},
		{opt: "pascal", word: "foo_bar", want: "FooBar"},
		{opt: "l", word: "foo_bar", want: "foo-bar"},
		{opt: "lisp", word: "foo_bar", want: "foo-bar"},
	}

	for _, tt := range tests {
		got := ChangeCase(tt.opt, tt.word)
		if got != tt.want {
			t.Errorf("ChangeCase(%q, %q) was %q, want %q",
				tt.opt, tt.word, got, tt.want)
		}
	}
}

func TestProcess(t *testing.T) {
	input := `fooBar this-is-a-test  snake_case
		ClassName`
	want := `foo_bar
this_is_a_test
snake_case
class_name
`

	in := strings.NewReader(input)
	var out bytes.Buffer
	if err := Process(in, &out, "snake"); err != nil {
		t.Fatalf("Process returned error: %v", err)
	}
	got := out.String()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
