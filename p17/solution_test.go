package p17

import (
	"reflect"
	"testing"
)

func TestExample(t *testing.T) {
	actual := letterCombinations("2")
	expected := []string{"a", "b", "c"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
		t.Fail()
	}
}

func TestExample1(t *testing.T) {
	actual := letterCombinations("23")
	expected := []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	actual := letterCombinations("29")
	expected := []string{"aw", "ax", "ay", "az", "bw", "bx", "by", "bz", "cw", "cx", "cy", "cz"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
		t.Fail()
	}
}
