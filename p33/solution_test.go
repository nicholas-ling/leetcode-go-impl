package p33

import (
	"testing"
)

func TestExample(t *testing.T) {
	actual := search([]int{0}, 9)
	if actual != -1 {
		t.Fail()
	}
}

func TestExample1(t *testing.T) {
	actual := search([]int{0}, 0)
	if actual != 0 {
		t.Log(actual)
		t.Fail()
	}
}

func TestExample2(t *testing.T) {
	actual := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	if actual != 4 {
		t.Log(actual)
		t.Fail()
	}
}

func TestExample3(t *testing.T) {
	actual := search([]int{4, 5, 6, 7, 0, 1, 2}, 5)
	if actual != 1 {
		t.Log(actual)
		t.Fail()
	}
}

func TestExample4(t *testing.T) {
	actual := search([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	if actual != -1 {
		t.Log(actual)
		t.Fail()
	}
}
