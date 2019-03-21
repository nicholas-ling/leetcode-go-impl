package p173

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	ordered []*TreeNode
	current int
}

func Constructor(root *TreeNode) BSTIterator {
	ret := new(BSTIterator)
	inorder(root, ret)
	return *ret
}

func inorder(root *TreeNode, it *BSTIterator) {
	if root == nil {
		return
	}
	inorder(root.Left, it)
	it.ordered = append(it.ordered, root)
	inorder(root.Right, it)
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	current := this.current
	this.current += 1
	return this.ordered[current].Val
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.current < len(this.ordered)
}
