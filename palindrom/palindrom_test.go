package palindrom

import "testing"

func TestCheckPolindrom(t *testing.T) {
	if !CheckPalindrome(131) {
		t.Error("it's not palidrom")
	}
	if CheckPalindrome(13321) {
		t.Error("itpalidrom")
	}
}
