package utils

/*
	Implements a tree using generics in go
*/

type orderable interface {
	~int8 | ~int16 | ~int32 | ~float32 | ~float64
}

type Node[T orderable] struct {
	val      T
	children []Node[T]
}

func (n Node[T]) insert(val T) {
	child := &Node[T]{val: val}
	n.children = append(n.children, *child)
}

type Tree[T orderable] []Node[T]
