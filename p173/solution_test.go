package p173

import (
	"testing"
)

func TestExample(t *testing.T) {
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 9}
	root.Right.Right = &TreeNode{Val: 20}
	obj := Constructor(root)
	if !obj.HasNext() || obj.Next() != 3 {
		t.Fail()
	}
	if !obj.HasNext() || obj.Next() != 7 {
		t.Fail()
	}
	if !obj.HasNext() || obj.Next() != 9 {
		t.Fail()
	}
	if !obj.HasNext() || obj.Next() != 15 {
		t.Fail()
	}
	if !obj.HasNext() || obj.Next() != 20 {
		t.Fail()
	}
	if obj.HasNext() {
		t.Fail()
	}
}

func TestEmpty(t *testing.T) {
	obj := Constructor(nil)
	if obj.HasNext() {
		t.Fail()
	}
}
