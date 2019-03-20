package p103

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root == nil {
		return ret
	}
	ls, rs := []*TreeNode{}, []*TreeNode{}
	ls = append(ls, root)

	for len(ls) > 0 || len(rs) > 0 {
		row := []int{}
		if len(ls) > 0 {
			for len(ls) > 0 {
				node := ls[len(ls)-1]
				ls = ls[:len(ls)-1]
				row = append(row, node.Val)
				if node.Left != nil {
					rs = append(rs, node.Left)
				}
				if node.Right != nil {
					rs = append(rs, node.Right)
				}
			}
		} else {
			for len(rs) > 0 {
				node := rs[len(rs)-1]
				rs = rs[:len(rs)-1]
				row = append(row, node.Val)
				if node.Right != nil {
					ls = append(ls, node.Right)
				}
				if node.Left != nil {
					ls = append(ls, node.Left)
				}
			}
		}
		ret = append(ret, row)
	}

	return ret
}
