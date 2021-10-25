## 二叉树
### 层次遍历（BFS）

https://github.com/RiKaCC/Go-DataStructure/blob/master/tree/bTreeBFS.go

```
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}

func bTreeBFS(root *TreeNode) (ret []int) {
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
```

### 先序遍历
```
func tree2str(t *TreeNode) []int {
	st := []*TreeNode{}
	st = append(st, t)
	nums := []int{}

	for len(st) > 0 {
		len := len(st)
		// 使用数组模拟栈，把最后一个元素（后入先出）取出来
		cur := st[len-1]
		nums = append(nums, cur.Val)
		// st = s[:len-1]模拟了栈的pop操作
		st = st[:len-1]
		if cur.Right != nil {
			st = append(st, cur.Right)
		}

		if cur.Left != nil {
			st = append(st, cur.Left)
		}
	}

	return nums
}
```

### 中序遍历

```
因为第一个结点不是root,所以需要在循环里进行压栈。
若当前结点不为空，将当前结点压栈，当前结点赋值为当前结点的左子结点，一直循环这一步。
若当前结点为空，则出栈，并置当前结点为出栈结点的右子树。

func findTarget(root *TreeNode, k int) []int {
	st := []*TreeNode{}
	nums := []int{}

	for len(st) > 0 || root != nil {
		len := len(st)
		if root != nil {
			st = append(st, root)
			root = root.Left
		} else {
			cur := st[len-1]
			nums = append(nums, cur.Val)
			st = st[:len-1]
			root = cur.Right
		}
	}

	return nums
}

// 递归的做法
func traverse(rooot *TreeNode) {
	// 前序
	traverse(root.Left)
	// 中序
	traverse(root.Right)
	// 后序
}
```
