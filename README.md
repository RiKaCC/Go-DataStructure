# Go-DataStructure
- [queue](#queue)
- [stack](#stack)
- [graph](#graph)
    - [bfs](#bfs)
    - [dijkstra](#dijkstra)
    - [prim](#prim)
- [BTree](https://github.com/RiKaCC/Go-DataStructure/tree/master/tree)

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
before make a graph, we need make node, edge.

the `node` datastructure like this:

```go
type node struct {
	id string
}
```
the `edge` datastructure like this:

it's easy to know each param means
```go
type Edge struct {
	src    Node
	dst    Node
	weight int
}
```
the `graph` datastructure like this:
```go
type Graph struct {
	sync.RWMutex
	edge    map[string]map[string]int
	nodeMap map[string]Node // record all the node in a graph
}
```
if node `A` to node `B`, the distance is 5,we can describe like this:
```go
Graph.dege[A][B] = 5
Graph.nodeMap[A] = nodeA
Graph.nodeMap[B] = nodeB
```

how we create a graph?
```go
g := NewGraph()
g.AddEdge("A", "B", 3) // A->B 3
g.AddEdge("A", "C", 2) // A->C 2
g.AddEdge("B", "E", 6) // B->E 6
g.AddEdge("B", "D", 2) // B->D 2
g.AddEdge("D", "E", 3) // D->E 3
g.AddEdge("C", "F", 1) // C-F 1
g.AddEdge("E", "F", 4) // E->F 4
```

### bfs
BFS(广度优先搜索)，是一种很常见的搜索。适用于无权最短路径的问题。

BFS的本质其实还是贪心算法。实现的过程中，需要借助queue。

我们还需要设置一个visited的数组（或map）来标记节点是否已被访问过，减少访问次数。

还需要设置一个记录最短距离的数组distance（或map）来记录最短距离。

1.将第一个结点放入queue,按照设定的规则搜索它周围的符合条件的结点

2.并将符合条件的结点放入queue，并设置visited为已访问，更新距离数组distance +1

3.遍历queue

### dijkstra
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

完整代码：
```golang
package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"math"
	"sync"
)

type node struct {
	id string
}

type Edge struct {
	src    Node
	dst    Node
	weight float64
}

type Graph struct {
	sync.RWMutex
	edge    map[string]map[string]float64
	nodeMap map[string]Node // record all the node in a graph
}

type Node interface {
	NodeID() string
}

func NewNode(id string) Node {
	return &node{id: id}
}

func (n *node) NodeID() string {
	return n.id
}

func NewEdge(src Node, dst Node, w float64) *Edge {
	return &Edge{src: src, dst: dst, weight: w}
}

func NewGraph() *Graph {
	return &Graph{
		edge:    make(map[string]map[string]float64),
		nodeMap: make(map[string]Node),
	}
}

func (g *Graph) AddEdge(nodeID1 string, nodeID2 string, w float64) {
	g.Lock()
	defer g.Unlock()

	if nodeID1 == nodeID2 {
		panic("can't add same vertex in one edge")
		return
	}

	if w == 0 {
		panic("weight can't use 0")
		return
	}

	// record each vertex
	g.nodeMap[nodeID1] = NewNode(nodeID1)
	g.nodeMap[nodeID2] = NewNode(nodeID2)

	if _, ok := g.edge[nodeID1]; ok {
		g.edge[nodeID1][nodeID2] = w
	} else {
		tempMap := make(map[string]float64)
		tempMap[nodeID2] = w
		g.edge[nodeID1] = tempMap
	}
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B", 3)
	g.AddEdge("A", "C", 2)
	g.AddEdge("B", "E", 5)
	g.AddEdge("B", "D", 2)
	g.AddEdge("D", "E", -2)
	g.AddEdge("C", "F", 1)
	g.AddEdge("E", "F", 3)

	fmt.Println(g)
	shortDis := g.Dijkstra("A", "E")
	fmt.Println("A->E shortest distance is:", shortDis)
}

func (g *Graph) Dijkstra(src string, dst string) (shortDis float64) {
	// infinity
	infinity := math.Inf(1)

	// init a short distance map to record the shortest distance from src
	distance := make(map[string]float64)
	for nodeID := range g.nodeMap {
		if nodeID == src {
			distance[nodeID] = 0
		} else {
			distance[nodeID] = infinity
		}
	}

	q := NewQueue()
	q.Push(src)

	for !q.Empty() {
		v := q.Pop()
		e, ok := v.(string)
		if !ok {
			return 0
		}

		for nodeID := range g.nodeMap {
			if nodeID == e {
				continue
			}

			if g.edge[e][nodeID]+distance[e] < distance[nodeID] && g.edge[e][nodeID] != 0 {
				distance[nodeID] = g.edge[e][nodeID] + distance[e]
				q.Push(nodeID)
			}
		}
	}

	for node := range g.nodeMap {
		temp := fmt.Sprintf("from A -> %s =  %d", node, int(distance[node]))
		fmt.Println(temp)
	}
	return distance[dst]
}
```

### prim
Prim算法思想其实比较简洁，首先需要去理解什么是最小生成树？下面是来自维基百科的解释：
> 最小生成树是一副连通加权无向图中一棵权值最小的生成树。
在一给定的无向图 G = (V, E) 中，(u, v) 代表连接顶点 u 与顶点 v 的边（即 {\displaystyle (u,v)\in E} (u,v)\in E），而 w(u, v) 代表此边的权重，若存在 T 为 E 的子集（即 {\displaystyle T\subseteq E} T\subseteq E）且 (V, T) 为树，使得
{\displaystyle w(T)=\sum _{(u,v)\in T}w(u,v)} w(T)=\sum _{(u,v)\in T}w(u,v)
的 w(T) 最小，则此 T 为 G 的最小生成树。
最小生成树其实是最小权重生成树的简称.

简单一点理解就是：从一个图中选出节点，添加到一个新树上的过程

**实现：**

1.将要处理的无向图看做未访问过的集合unknown

2.定义一个新的集合来记录已经访问的节点known

3.将初始节点放入known中

4.遍历unknown,从中找到路基known中节点权值最小的点，将该点加入known,并从unknown中剔除

5.执行4步骤，直到unknown中没有节点，这样生成的最小生成树就是known

**Implement**

1.at first, wo should make the graph like a unknown set

2.difine a new set named known to record the node which visited

3.put the src node into knowd set

4.find a node from unknown which not in known , also has the shortest distance from the nodes in knowd

5.repeate step 4

完整代码如下：
https://github.com/RiKaCC/Go-DataStructure/blob/master/Prim.go

```golang
package main

import (
	"fmt"
	//"github.com/davecgh/go-spew/spew"
	"sync"
)

type node struct {
	id string
}

type Edge struct {
	src    Node
	dst    Node
	weight float64
}

type Graph struct {
	sync.RWMutex
	edge    map[string]map[string]float64
	nodeMap map[string]Node // record all the node in a graph
}

type Node interface {
	NodeID() string
}

func NewNode(id string) Node {
	return &node{id: id}
}

func (n *node) NodeID() string {
	return n.id
}

func NewEdge(src Node, dst Node, w float64) *Edge {
	return &Edge{src: src, dst: dst, weight: w}
}

func NewGraph() *Graph {
	return &Graph{
		edge:    make(map[string]map[string]float64),
		nodeMap: make(map[string]Node),
	}
}

func (g *Graph) AddEdge(nodeID1 string, nodeID2 string, w float64) {
	g.Lock()
	defer g.Unlock()

	if nodeID1 == nodeID2 {
		panic("can't add same vertex in one edge")
		return
	}

	if w == 0 {
		panic("weight can't use 0")
		return
	}

	// record each vertex
	g.nodeMap[nodeID1] = NewNode(nodeID1)
	g.nodeMap[nodeID2] = NewNode(nodeID2)

	if _, ok := g.edge[nodeID1]; ok {
		g.edge[nodeID1][nodeID2] = w
	} else {
		tempMap := make(map[string]float64)
		tempMap[nodeID2] = w
		g.edge[nodeID1] = tempMap
	}
}

func main() {
	g := NewGraph()
	g.AddEdge("A", "B", 7)
	g.AddEdge("A", "D", 5)
	g.AddEdge("B", "C", 8)
	g.AddEdge("B", "A", 7)
	g.AddEdge("B", "D", 9)
	g.AddEdge("B", "E", 7)
	g.AddEdge("C", "B", 8)
	g.AddEdge("C", "E", 5)
	g.AddEdge("D", "A", 5)
	g.AddEdge("D", "B", 9)
	g.AddEdge("D", "F", 6)
	g.AddEdge("D", "E", 15)
	g.AddEdge("D", "F", 6)
	g.AddEdge("E", "B", 7)
	g.AddEdge("E", "C", 5)
	g.AddEdge("E", "D", 15)
	g.AddEdge("E", "F", 8)
	g.AddEdge("E", "G", 9)
	g.AddEdge("F", "D", 6)
	g.AddEdge("F", "E", 8)
	g.AddEdge("F", "G", 11)
	g.AddEdge("G", "E", 9)
	g.AddEdge("G", "F", 11)

	fmt.Println(g)
	g.Prim("D")
}

func (g *Graph) Prim(src string) {
	knowNode := make(map[int]string)
	unknowNode := g.nodeMap
	tempEdgeMap := g.edge
	edgeMap := make(map[int]*Edge)
	totalWeight := float64(0)
	key := 0

	knowNode[key] = src
	delete(unknowNode, knowNode[key])
	key++
	var temp string

	// 找到与knowNodeMap里节点权值最小的节点，并将该节点加入nodeMap
	for len(unknowNode) > 0 {
		min := float64(1000000)
		for nodeID := range unknowNode {
			for _, v := range knowNode {
				if tempEdgeMap[nodeID][v] < min && tempEdgeMap[nodeID][v] != 0 {
					min = tempEdgeMap[nodeID][v]
					knowNode[key] = nodeID
					temp = v
				}
			}
		}

		n1 := NewNode(temp)
		n2 := NewNode(knowNode[key])
		edge := NewEdge(n1, n2, min)
		edgeMap[key-1] = edge

		// 从未知节点map删除已经找到的当前权值最小的节点
		delete(unknowNode, knowNode[key])
		totalWeight += min
		key++
	}

	fmt.Println(totalWeight)
	fmt.Println(knowNode)
}
```
