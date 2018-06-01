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
	g.AddEdge("B", "E", 6)
	g.AddEdge("B", "D", 2)
	g.AddEdge("D", "E", 3)
	g.AddEdge("C", "F", 1)
	g.AddEdge("E", "F", 3)

	fmt.Println(g)
	shortDis := g.Dijkstra("A", "E")
	fmt.Println("A->F shortest distance is:", shortDis)
}

func (g *Graph) Dijkstra(src string, dst string) (shortDis float64) {
	// infinity
	infinity := math.Inf(1)

	// at first we need init a map to record the distance of every two vertex
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

			if g.edge[e][nodeID]+distance[e] < distance[nodeID] && g.edge[e][nodeID] > 0 {
				distance[nodeID] = g.edge[e][nodeID] + distance[e]
				q.Push(nodeID)
			}
		}
	}
	fmt.Println(distance)
	return distance[dst]
}
