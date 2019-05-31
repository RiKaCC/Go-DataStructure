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
		cur := st[len-1]
		nums = append(nums, cur.Val)
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
