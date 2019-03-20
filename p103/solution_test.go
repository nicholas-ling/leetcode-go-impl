package p103

import (
	"reflect"
	"testing"
)

func TestExample(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 7}

	actual := zigzagLevelOrder(root)
	expected := [][]int{[]int{1}, []int{2, 3}, []int{4, 5, 6, 7}}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
		t.Fail()
	}
}

func TestEmpty(t *testing.T) {
	actual := zigzagLevelOrder(nil)
	expected := [][]int{}
	if !reflect.DeepEqual(actual, expected) {
		t.Error(actual, expected)
		t.Fail()
	}
}
