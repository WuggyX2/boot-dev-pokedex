package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "test",
			expected: []string{"test"},
		},
		{
			input:    "TEST",
			expected: []string{"test"},
		},
		{
			input:    "",
			expected: []string{},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Result slice size (%d) does not match expected size %d", len(actual), len(c.expected))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual word %s did not match expected word %s", word, expectedWord)
			}
		}
	}
}
