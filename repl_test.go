package main

import (
    "testing"
)

func TestCleanInput(t *testing.T) {
    cases := []struct {
        input    string
        expected []string
    }{
        {
            input:    "  hello world ",
            expected: []string{"hello", "world"},
        },
        {
            input:    " who's that pokemon ",
            expected: []string{"who's", "that", "pokemon"},
        },
        {
            input:    " diglett DUGtrio ",
            expected: []string{"diglett", "dugtrio"},
        },
    }
    for _, c := range cases {
        actual := cleanInput(c.input)
        if len(actual) != len(c.expected) {
            t.Errorf("Length of slice not expected.")
        }
        for i := range actual {
            word := actual[i]
            expectedWord := c.expected[i]
            if word != expectedWord {
                t.Errorf("Word does not match expected word.")
            }
        }
    }
}
