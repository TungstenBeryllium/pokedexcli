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
			input: "pikachu gorochu",
			expected: []string{"pikachu", "gorochu"},

		},

		{
			input: "Goro Majima and Kotomine Kirei",
			expected: []string{"goro", "majima", "and", "kotomine", "kirei"},

		},
	
	}


	for _, c := range cases {
		actual := cleanInput(c.input)
	// Check the length of the actual slice
	// if they don't match, use t.Errorf and continue to the next case
		if len(actual) != len(c.expected) {
			t.Errorf("error occured")
			continue
		}
	
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			
			if word != expectedWord {
				t.Errorf("error occured")
			}
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test

			
		}

	
	}
}
