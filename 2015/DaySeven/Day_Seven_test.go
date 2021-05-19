package dayseven

import "testing"

func TestLitLightsOne(t *testing.T) {
	input := []string{"123 -> x", "456 -> y", "x AND y -> d","x OR y -> e","x LSHIFT 2 -> f","y RSHIFT 2 -> g","NOT x -> h","NOT y -> i"}
	results := ApplyOperations(input)

	if results["d"] != 72 {
		t.Errorf("d should have been 72, but was: %v", results["d"])
	}
	if results["e"] != 507 {
		t.Errorf("e should have been 507, but was: %v", results["e"])
	}
	if results["f"] != 492 {
		t.Errorf("f should have been 492, but was: %v", results["f"])
	}
	if results["g"] != 114 {
		t.Errorf("g should have been 114, but was: %v", results["g"])
	}
	if results["h"] != 65412 {
		t.Errorf("h should have been 65412, but was: %v", results["h"])
	}
	if results["i"] != 65079 {
		t.Errorf("i should have been 65079, but was: %v", results["i"])
	}
	if results["x"] != 123 {
		t.Errorf("x should have been 123, but was: %v", results["x"])
	}
	if results["y"] != 456 {
		t.Errorf("y should have been 456, but was: %v", results["y"])
	}
}