package utils

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGenerateHumanID(t *testing.T) {
	// Test case 1: Empty prefix
	expectedPrefix := "LGM"
	expectedNow := time.Now()
	expectedNs := expectedNow.UnixMilli() % millisecondsPerSecond

	expectedID := fmt.Sprintf("%s%d-%03d-%02d%02d%02d%03d", strings.ToUpper(expectedPrefix), expectedNow.Year()-2000, expectedNow.YearDay(), expectedNow.Hour(), expectedNow.Minute(), expectedNow.Second(), expectedNs)

	result := GenerateHumanID("")
	if result != expectedID {
		t.Errorf("GenerateHumanID with empty prefix returned unexpected ID: got %v, want %v", result, expectedID)
	}

	// Test case 2: Non-empty prefix
	customPrefix := "ABC"
	expectedID = fmt.Sprintf("%s%d-%03d-%02d%02d%02d%03d", strings.ToUpper(customPrefix), expectedNow.Year()-2000, expectedNow.YearDay(), expectedNow.Hour(), expectedNow.Minute(), expectedNow.Second(), expectedNs)

	result = GenerateHumanID(customPrefix)
	if result != expectedID {
		t.Errorf("GenerateHumanID with custom prefix returned unexpected ID: got %v, want %v", result, expectedID)
	}
}
