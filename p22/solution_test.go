package p22

import (
	"reflect"
	"testing"
)

func TestExample1(t *testing.T) {
	actual := generateParenthesis(1)
	expected := []string{"()"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	actual := generateParenthesis(2)
	expected := []string{"(())", "()()"}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
		t.Fail()
	}
}
