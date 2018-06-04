[TOC]
# Go-DataStructure
## English
On the way of learning Golang.

Go-DataStructure is implement the normal datastructure on golang.

### queue
queue implement simple 3 method of queue.`Pop`, `Push`, `Empty`

#### Push()
you can push any type as a element into queue.
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

## 中文
学习golang的路上。

Go-DataStructure基于golang来实现了部分常用数据结构。


