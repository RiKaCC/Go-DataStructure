# Go-DataStructure
- [queue](#queue)
- [stack](#stack)
- [BFS](#BFS)
- [Dijkstra](#Dijkstra)

## English
On the way of learning Golang.

Go-DataStructure is implement the normal datastructure on golang.

### queue
queue implement simple 3 methods of queue.`Pop`, `Push`, `Empty`

#### Push()
you can push any type element into a queue.
```golang
q := NewQueue()
e := 2
q.Push(e)
```
#### Pop()
```golang 
e := q.Pop()
fmt.Println(e)
```
#### Empty()
judgement the queue is empty. if the queue is empty, Empty() will return true
```golang
if q.Empty() {
  // the queue is empty
}
```
### stack
stack implement simple 4 methods of stack. `Push`, `Pop`, `Empty`, `SetEmpty`.

the stack struct simple defintion as flow:
```golang
type Stack struct {
  data []interface{}
  size int
}
```

#### Push()
you can push any type element into a stack.
```golang
s := NewStack()
e1 := 1
e2 := "hello world"
s.Push(e1)
s.Push(e2)
```
#### Pop
```
e := s.Pop()
fmt.Println(e)
```

#### Empty(), SetEmpty()
judgement the stack is empty, if stack is empty Empty() will return true.

SetEmpty() set the stack empty
```golang
if s.Empty() {
  // the stack is empty
} else {
  s.SetEmpty()
}
```

## 中文
学习golang的路上。

Go-DataStructure基于golang来实现了部分常用数据结构。


