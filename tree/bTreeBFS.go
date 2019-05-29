package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//       tree
//        1
//      2   3
//    4  5 6  7
//
// BFS out put : 1 2 3 4 5 6 7

func BTreeBFS(root *TreeNode) (ret []int) {
	// 用数组模拟队列
	que := []*TreeNode{}
	que = append(que, root)

	for len(que) > 0 {
		s := len(que) // 用来控制每一层的遍历
		for i := 0; i < s; i++ {
			cur := que[i]
			ret = append(ret, cur.Val)

			if cur.Left != nil {
				que = append(que, cur.Left)
			}

			if cur.Right != nil {
				que = append(que, cur.Right)
			}
		}

		que = que[s:] // 弹出已经取出的节点
	}

	return ret
}

func main() {
	node7 := &TreeNode{Val: 7, Left: nil, Right: nil}
	node6 := &TreeNode{Val: 6, Left: nil, Right: nil}
	node5 := &TreeNode{Val: 5, Left: nil, Right: nil}
	node4 := &TreeNode{Val: 4, Left: nil, Right: nil}
	node3 := &TreeNode{Val: 3, Left: node6, Right: node7}
	node2 := &TreeNode{Val: 2, Left: node4, Right: node5}
	node1 := &TreeNode{Val: 1, Left: node2, Right: node3}

	fmt.Println(BTreeBFS(node1))
}
