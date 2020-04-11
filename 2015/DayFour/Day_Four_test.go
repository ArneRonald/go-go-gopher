package dayone

import (
	"testing"
)

func TestScenarioOne(t *testing.T) {
	result := findInput("abcdef")
	if result != 609043 {
		t.Errorf("Expected 609043, Found %v", result)
	}
}

func TestScenarioTwo(t *testing.T) {
	result := findInput("pqrstuv")
	if result != 1048970 {
		t.Errorf("Expected 1048970, Found %v", result)
	}
}
