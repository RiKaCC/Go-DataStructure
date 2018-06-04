# Go-DataStructure
- [queue](#queue)
- [stack](#stack)
- [graph](#graph)
  - [BFS](#BFS)
  - [Dijkstra](#Dijkstra)

On the way of learning Golang.

Go-DataStructure is implement the normal datastructure on golang.

## queue
queue implement simple 3 methods of queue.`Pop`, `Push`, `Empty`

### Push()
you can push any type element into a queue.
```golang
q := NewQueue()
e := 2
q.Push(e)
```
### Pop()
```golang 
e := q.Pop()
fmt.Println(e)
```
### Empty()
judgement the queue is empty. if the queue is empty, Empty() will return true
```golang
if q.Empty() {
  // the queue is empty
}
```
## stack
stack implement simple 4 methods of stack. `Push`, `Pop`, `Empty`, `SetEmpty`.

the stack struct simple defintion as flow:
```golang
type Stack struct {
  data []interface{}
  size int
}
```

### Push()
you can push any type element into a stack.
```golang
s := NewStack()
e1 := 1
e2 := "hello world"
s.Push(e1)
s.Push(e2)
```
### Pop
```
e := s.Pop()
fmt.Println(e)
```

### Empty(), SetEmpty()
judgement the stack is empty, if stack is empty Empty() will return true.

SetEmpty() set the stack empty
```golang
if s.Empty() {
  // the stack is empty
} else {
  s.SetEmpty()
}
```

## graph

### BFS
BFS(广度优先搜索)，是一种很常见的搜索。适用于无权最短路径的问题。

BFS的本质其实还是贪心算法。实现的过程中，需要借助queue。

我们还需要设置一个visited的数组（或map）来标记节点是否已被访问过，减少访问次数。

还需要设置一个记录最短距离的数组distance（或map）来记录最短距离。

1.将第一个结点放入queue,按照设定的规则搜索它周围的符合条件的结点

2.并将符合条件的结点放入queue，并设置visited为已访问，更新距离数组distance +1

3.遍历queue

### Dijkstra
BFS能够解决无权最短路径问题，但是如果路径加权了，BFS就没法解决了，这时候就需要使用Dijkstra来解决这个问题。

Dijkstra的答题思路跟BFS差不多，举一反三。

首先同样也需要一个记录最短距离的数组distance, 我们得先初始化这个数组。初始化距离都为整无穷大。伪代码如下：

遍历整个图，挨个给结点v初始化最短距离为正无穷大，如果是起始结点，那么最短距离就是0
```golang
distance := make(map[string]float64)

for v := range graph {
    if v = src {
        distance[v] = 0
    } else {
        distance[v] = infinity
    }
}
```

