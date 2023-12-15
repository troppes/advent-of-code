package util

import "fmt"

type Node[T any] struct {
	Info T
	Next *Node[T]
}

type List[T any] struct {
	Head  *Node[T]
	match func(*T, *T) bool // Custom function for matching
}

func (l *List[T]) SetMatcher(t func(*T, *T) bool) {
	l.match = t
}

func (l *List[T]) Insert(d T) {
	list := &Node[T]{Info: d, Next: nil}
	if l.Head == nil {
		l.Head = list
	} else {
		p := l.Head
		for p.Next != nil {
			p = p.Next
		}
		p.Next = list
	}
}

func (l *List[T]) Print() {
	p := l.Head
	for p != nil {
		fmt.Printf("-> %v ", p.Info)
		p = p.Next
	}
	fmt.Println()
}

func (l *List[T]) Empty() bool {
	return l.Head == nil
}

func (l *List[T]) InsertAndReplace(d T) {
	item := &Node[T]{Info: d, Next: nil}
	if l.Head == nil {
		l.Head = item
		return
	}

	curr := l.Head
	if l.match(&curr.Info, &d) { // check for head
		curr.Info = d
		return
	}

	for curr.Next != nil {
		if l.match(&curr.Next.Info, &d) {
			curr.Next.Info = d
			return // Found and replaced; no need to continue
		}
		curr = curr.Next
	}

	// If the loop ends without finding a match, insert the new node at the end
	curr.Next = item
}

func (l *List[T]) Delete(data T) {

	curr := l.Head
	if curr == nil {
		return
	}

	if l.match(&curr.Info, &data) {
		l.Head = curr.Next
		return
	}

	// Handling deletion in the middle or end of the list
	for curr.Next != nil {
		if l.match(&curr.Next.Info, &data) {
			if curr.Next.Next != nil {
				curr.Next = curr.Next.Next
			} else {
				curr.Next = nil // Deleting the last node
			}
			return
		}
		curr = curr.Next
	}
}
