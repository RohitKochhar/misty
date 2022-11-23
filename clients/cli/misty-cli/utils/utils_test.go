package utils_test

import (
	"fmt"
	"rohitsingh/misty-cli/utils"
	"testing"
)

func TestSanitizeTopic(t *testing.T) {
	// Using table-driven testing to cover
	// a variety of potential inputs
	testCases := []struct {
		name   string // Name of test case
		input  string // Input to SanitizeTopic function
		output string // Expected output from SanitizeTopic function
	}{
		{"AlreadySanitizedTopic", "/topic", "/topic"},
		{"NoLeadingDash", "topic", "/topic"},
		{"LeadingAndTrailingDash", "/topic/", "/topic"},
		{"NoLeadingDashButTrailingDash", "topic/", "/topic"},
	}
	for _, tc := range testCases {
		testName := fmt.Sprintf("Test%s", tc.name)
		t.Run(testName, func(t *testing.T) {
			out, err := utils.SanitizeTopic(tc.input)
			if err != nil {
				t.Fatalf("unexpected error while sanitizing %s: %q", tc.input, err)
			}
			if out != tc.output {
				t.Fatalf("expected %s, instead got %s", tc.output, out)
			}
		})
	}
}
