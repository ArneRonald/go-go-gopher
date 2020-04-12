package dayfive

import "testing"

func TestContainsThreeVowels(t *testing.T) {
	if !containsThreeVowels("aei") {
		t.Errorf("Expected result as true but returned false")
	}
	if !containsThreeVowels("xazegov") {
		t.Errorf("Expected result as true but returned false")
	}
	if !containsThreeVowels("aeiouaeiouaeiou") {
		t.Errorf("Expected result as true but returned false")
	}
	if containsThreeVowels("plwrtnbms") {
		t.Errorf("Expected result as false but returned true")
	}
}

func TestContainsOneLetterThatAppearsTwiceInARow(t *testing.T) {
	if !containsOneLetterThatAppearsTwiceInARow("xx") {
		t.Errorf("Expected result as true but returned false")
	}
	if !containsOneLetterThatAppearsTwiceInARow("abcdde") {
		t.Errorf("Expected result as true but returned false")
	}
	if !containsOneLetterThatAppearsTwiceInARow("aabbccdd") {
		t.Errorf("Expected result as true but returned false")
	}
	if containsOneLetterThatAppearsTwiceInARow("abcde") {
		t.Errorf("Expected result as false but returned true")
	}
}

func TestContainsRestrictedWords(t *testing.T) {
	if !containRestrictedWords("ab") {
		t.Errorf("Expected true but returned false")
	}
	if !containRestrictedWords("cd") {
		t.Errorf("Expected true but returned false")
	}
	if !containRestrictedWords("pq") {
		t.Errorf("Expected true but returned false")
	}
	if !containRestrictedWords("xy") {
		t.Errorf("Expected true but returned false")
	}
}

func TestDetermineNaughtyOrNice(t *testing.T) {
	if !determineNaughtyOrNice("ugknbfddgicrmopn") {
		t.Errorf("Name conforms to nice but was returned as Naughty")
	}
	if !determineNaughtyOrNice("aaa") {
		t.Errorf("Name conforms to nice but was returned as Naughty")
	}
	if determineNaughtyOrNice("jchzalrnumimnmhp") {
		t.Errorf("Name conforms to Naughty but was returned as Nice")
	}
	if determineNaughtyOrNice("haegwjzuvuyypxyu") {
		t.Errorf("Name conforms to Naughty but was returned as Nice")
	}
	if determineNaughtyOrNice("dvszwmarrgswjxmb") {
		t.Errorf("Name conforms to Naughty but was returned as Nice")
	}
}

func TestDetermineNaughtyOrNiceNew(t *testing.T) {
	if !determineNaughtyOrNiceNew("qjhvhtzxzqqjkmpb") {
		t.Errorf("Name conforms to nice but was returned as Naughty")
	}
	if !determineNaughtyOrNiceNew("xxyxx") {
		t.Errorf("Name conforms to nice but was returned as Naughty")
	}
	if determineNaughtyOrNiceNew("uurcxstgmygtbstg") {
		t.Errorf("Name conforms to naughty but was returned as nice")
	}
	if determineNaughtyOrNiceNew("ieodomkazucvgmuy") {
		t.Errorf("Name conforms to naughty but was returned as nice")
	}
}

func TestContainsPairOfTwoLetters(t *testing.T) {
	if !containsPairOfTwoLetters("xyxy") {
		t.Errorf("Expected true but returned false")
	}
	if !containsPairOfTwoLetters("aabcdefgaa") {
		t.Errorf("Expected true but returned false")
	}
	if containsPairOfTwoLetters("aaa") {
		t.Errorf("Expected false but returned true")
	}
}

func TestContainsRepeatingLettersWithOneInbetween(t *testing.T) {
	if !containsRepeatingLettersWithOneInbetween("xyx") {
		t.Errorf("Expected true but returned false")
	}
	if !containsRepeatingLettersWithOneInbetween("abcdefeghi") {
		t.Errorf("Expected true but returned false")
	}
	if !containsRepeatingLettersWithOneInbetween("aaa") {
		t.Errorf("Expected true but returned false")
	}
	if containsRepeatingLettersWithOneInbetween("rtyd") {
		t.Errorf("Expected false but returned true")
	}
}
