package p3

import (
	"testing"
)

func TestExample(t *testing.T) {
	actual := lengthOfLongestSubstring("abcabcbb")
	if actual != 3 {
		t.Fail()
	}
}

func TestExample1(t *testing.T) {
	actual := lengthOfLongestSubstring("bbbbb")
	if actual != 1 {
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	actual := lengthOfLongestSubstring("pwwkew")
	if actual != 3 {
		t.Fail()
	}
}

func TestExample3(t *testing.T) {
	actual := lengthOfLongestSubstring("abcad")
	if actual != 4 {
		t.Fail()
	}
}

func TestExample4(t *testing.T) {
	actual := lengthOfLongestSubstring("abba")
	if actual != 2 {
		t.Fail()
	}
}
func TestEmpty(t *testing.T) {
	actual := lengthOfLongestSubstring("")
	if actual != 0 {
		t.Fail()
	}
}
