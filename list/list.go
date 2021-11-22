package list

import (
	"strings"
	"fmt"
)

type node[T any] struct {
	val T
	next *node[T]
}

// List[T] is a singly-linked list of T
type List[T any] struct {
	head *node[T]
}

func (l List[T]) String() string {
	var buf strings.Builder

	buf.WriteRune('[')
	for n := l.head; n != nil; n = n.next {
		fmt.Fprintf(&buf, "%v", n.val)
		if n.next != nil {
			buf.WriteString(", ")
		}
	}
	buf.WriteRune(']')
	return buf.String()
}

// Make creates a list of T with a set of provided values. It's called Make
// instead of New because it performs O(n) allocations, where n == len(vals)
func Make[T any](vals ...T) List[T] {
	var l List[T]
	for i := len(vals)-1; i >= 0; i-- {
		l.Prepend(vals[i])
	}
	return l
}

// Empty is true for empty lists
func (l List[T]) Empty() bool {
	return l.head == nil
}

// Head returns the first element of the list. If the list is empty, Head
// returns the zero-value for the type T. This is the same thing as Peek()
func (l List[T]) Head() T {
	if l.head == nil {
		var v T
		return v
	}
	return l.head.val
}

// Pop returns the first element of the list and removes it from the list.
func (l *List[T]) Pop() T {
	if l.Empty() {
		var v T
		return v
	}

	v := l.head.val
	l.head = l.head.next
	return v
}

// Tail returns a list which is the original list without its Head element.
// If the original list is an empty list or a list of size 1, Tail is an
// empty list.
func (l List[T]) Tail() List[T] {
	if l.head == nil || l.head.next == nil {
		return List[T]{}
	}
	return List[T]{head: l.head.next}
}

// Len is the length of the list
func (l List[T]) Len() int {
	if l.head == nil {
		return 0
	}

	i := 0
	for  n := l.head; n != nil; n = n.next {
		i++
	}
	return i
}

// Prepend adds an element to the front of the list
func (l *List[T]) Prepend(v T) {
	l.head = &node[T]{
		val: v,
		next: l.head,
	}
}

// Map applies the input function f to each element of the list l, returning a
// new list containing the values produced by f
func (l List[T]) Map(f func(T) T) List[T] {
	var mapped List[T]

	if l.Empty() {
		return mapped
	}

	mapped.head = &node[T]{val: f(l.head.val)}
	last := mapped.head
	for n := l.head.next; n != nil; n = n.next {
		last.next = &node[T]{val: f(n.val)}
		last = last.next
	}

	return mapped
}
