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
	tmp := l.Footer
	if tmp == nil {
		tmp = &Node[T]{
			Node:  nil,
			Value: val,
		}
		l.Footer, l.Header = tmp, tmp
	} else if tmp.Node == nil {
		tmp.Node = &Node[T]{
			Node:  nil,
			Value: val,
		}
		l.Footer = tmp.Node
	}
}

func (l *List[T]) Pop() {
	var tmp *Node[T]
	if l.Header == nil {
		fmt.Println("No Element")
		goto x
	} else if l.Header.Node == nil {
		l.Header = nil
		l.Footer = nil
		fmt.Println("Only one Element")
		goto x
	} else {
		tmp = l.Header
	}

	for tmp.Node != nil {
		if tmp.Node == l.Footer {
			l.Footer = tmp
			tmp.Node = nil
			goto x
		}
		tmp = tmp.Node
	}
x:
	return
}

func (l *List[T]) Print() {
	var tmp = l.Header
	for tmp != nil {
		fmt.Print(tmp.Value, " ")
		tmp = tmp.Node
	}
	fmt.Println()
}
