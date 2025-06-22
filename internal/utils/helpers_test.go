package utils

import "testing"

func TestGenerateSlug(t *testing.T) {
	tests := []struct {
		input  string
		output string
		err    bool
	}{
		{"Hello World", "hello-world", false},
		{"  Hello   World  ", "hello-world", false},
		{"Hello & World", "hello-and-world", false},
		{"This -- is -- great!", "this-is-great", false},
		{"Go@Lang!", "golang", false},
		{"", "", true},
		{"   ", "", true},
		{"Clean-Up 123", "clean-up-123", false},
	}

	for _, test := range tests {
		slug, err := GenerateSlug(test.input)
		if test.err {
			if err == nil {
				t.Errorf("Expected error for input %q, but got none", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", test.input, err)
			}
			if slug != test.output {
				t.Errorf("Slugify(%q) = %q; want %q", test.input, slug, test.output)
			}
		}
	}
}
