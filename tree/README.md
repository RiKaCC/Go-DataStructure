## 二叉树
### 层次遍历（BFS）
```
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

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
```
