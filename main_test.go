package main

import (
	"testing"
)

func TestGetHelloWorld(t *testing.T) {
	var tests = []struct {
		name           string
		wantStatusCode int
		wantText       string
	}{
		{"default_200_response", 200, "Hello, World!"},
	}

	for _, test := range tests {
		gotStatusCode, gotText := getHelloWorld()
		if gotStatusCode != test.wantStatusCode || gotText != test.wantText {
			t.Errorf("[test: %s] got status: %d, text: %s. Wanted status: %d, text: %s",
				test.name,
				gotStatusCode,
				gotText,
				test.wantStatusCode,
				test.wantText,
			)
		}
	}
}
