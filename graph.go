package main

import (
	"fmt"
	"sync"
)

type Node struct {
	id int
}

type Edge struct {
	src    *Node
	dst    *Node
	weight int
}

type Graph struct {
	sync.RWMutex
	edge map[int]map[int]int
}

func NewNode(id int) *Node {
	return &Node{id: id}
}

func NewEdge(src *Node, dst *Node, w int) *Edge {
	return &Edge{src: src, dst: dst, weight: w}
}

func NewGraph() *Graph {
	return &Graph{edge: make(map[int]map[int]int)}
}

func (g *Graph) AddEdge(nodeID1 int, nodeID2 int, w int) {
	g.Lock()
	defer g.Unlock()

	if nodeID1 == nodeID2 {
		panic("can't add same vertex in one edge")
	}

	if _, ok := g.edge[nodeID1]; ok {
		g.edge[nodeID1][nodeID2] = w
	} else {
		tempMap := make(map[int]int)
		tempMap[nodeID2] = w
		g.edge[nodeID1] = tempMap
	}
}

func main() {
	g := NewGraph()
	g.AddEdge(1, 3, 2)
	g.AddEdge(1, 2, 4)
	g.AddEdge(1, 4, 15)
	g.AddEdge(2, 4, 6)
	g.AddEdge(2, 5, 3)
	g.AddEdge(3, 4, 7)
	g.AddEdge(3, 5, 1)
	g.AddEdge(4, 5, 2)

	fmt.Println(g.edge)
}
