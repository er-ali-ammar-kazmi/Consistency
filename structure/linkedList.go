package structure

import "fmt"

type List[T any] struct {
	Header *Node[T]
	Footer *Node[T]
}

type Node[T any] struct {
	Node  *Node[T]
	Value T
}

func NewList[T any]() List[T] {
	list := List[T]{}
	return list
}

func (l *List[T]) Push(val T) {
	tmp := l.Header
	for {
		if tmp == nil {
			tmp = &Node[T]{
				Node:  nil,
				Value: val,
			}
			l.Footer, l.Header = tmp, tmp
			return
		} else if tmp.Node == nil {
			tmp.Node = &Node[T]{
				Node:  nil,
				Value: val,
			}
			l.Footer = tmp.Node
			return
		}
		tmp = tmp.Node
	}
}

func (l *List[T]) Pop() {
	var last, tmp *Node[T]
	if l.Header == nil {
		fmt.Println("No Element")
		return
	} else if l.Header.Node == nil {
		l.Header = nil
		l.Footer = nil
		fmt.Println("Only one Element")
		return
	} else {
		last = l.Header
		tmp = l.Header.Node
	}

	for tmp.Node != nil {
		last = tmp
		tmp = tmp.Node
	}
	last.Node = nil
}

func (l *List[T]) Print() {
	var tmp *Node[T] = l.Header
	for tmp != nil {
		fmt.Print(tmp.Value, " ")
		tmp = tmp.Node
	}
	fmt.Println()
}
