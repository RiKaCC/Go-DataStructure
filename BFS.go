package main

import "fmt"

type Node struct {
	x int
	y int
}

func BFS(n *Node, b [][]int) int {
	q := NewQueue()
	q.Push(n)

	// 上下左右，四个方向
	dir := [][]int{{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0}}

	visit := make([][]int, 0)
	for i := 0; i < 10; i++ {
		temp := make([]int, 0)
		for j := 0; j < 10; j++ {
			temp = append(temp, 0)
		}
		visit = append(visit, temp)
	}

	max := 1
	for !q.Empty() {
		t := q.Pop()

		// 取出interface中的struct
		e, ok := t.(*Node)
		if !ok {
			return 0
		}

		visit[e.x][e.y] = 1
		for i := 0; i < 4; i++ {
			next_x := e.x + dir[i][0]
			next_y := e.y + dir[i][1]

			if next_x >= 0 && next_x < 5 && next_y >= 0 && next_y < 5 {
				if b[next_x][next_y] == 1 && visit[next_x][next_y] == 0 {
					next_n := &Node{
						x: next_x,
						y: next_y,
					}
					q.Push(next_n)
					visit[next_x][next_y] = 1
					max++
				}
			}
		}
	}
	return max
}

func main() {
	//	var a [][]int
	//	for i := 0; i < 10; i++ {
	//		s := make([]int, 0)
	//		for j := 0; j < 10; j++ {
	//			s = append(s, j)
	//		}
	//		a = append(a, s)
	//	}
	//

	max := 0
	//	fmt.Println(a)
	b := [][]int{{1, 0, 0, 1, 1},
		{0, 1, 0, 1, 0},
		{0, 0, 1, 1, 0},
		{1, 1, 0, 0, 1},
		{1, 0, 1, 0, 1}}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b[i][j] != 0 {
				n := &Node{
					x: i,
					y: j,
				}
				temp_max := BFS(n, b)
				if temp_max > max {
					max = temp_max
				}
			}
		}
	}
	fmt.Println("=================max:", max)
}
