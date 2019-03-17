package p103

import "fmt"
import "github.com/golang-collections/collections/stack"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 7}
	root.Left.Right = &TreeNode{Val: 6}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 4}
	fmt.Println(zigzagLevelOrder(root))
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	ls, rs := stack.New(), stack.New()
	ls.Push(root)

	for ls.Len() > 0 || rs.Len() > 0 {
		row := []int{}
		if ls.Len() > 0 {
			for ls.Len() > 0 {
				node := ls.Pop().(*TreeNode)
				row = append(row, node.Val)
				if node.Right != nil {
					rs.Push(node.Right)
				}
				if node.Left != nil {
					rs.Push(node.Left)
				}
			}
		} else {
			for rs.Len() > 0 {
				node := rs.Pop().(*TreeNode)
				row = append(row, node.Val)
				if node.Left != nil {
					ls.Push(node.Left)
				}
				if node.Right != nil {
					ls.Push(node.Right)
				}
			}

		}
		ret = append(ret, row)
	}

	return ret
}
