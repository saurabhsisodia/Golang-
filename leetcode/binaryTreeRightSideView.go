/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {

	var q []*TreeNode = []*TreeNode{}
	var ans []int = []int{}
	if root == nil {
		return ans
	}
	q = append(q, root)
	count := 1
	for len(q) != 0 {

		for count > 0 {
			n := q[0]
			q = q[1:]
			if count == 1 {
				ans = append(ans, n.Val)
			}
			count--
			if n.Left != nil {
				q = append(q, n.Left)
			}
			if n.Right != nil {
				q = append(q, n.Right)
			}

		}
		count = len(q)

	}
	return ans
}
