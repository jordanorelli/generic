package list

import (
	"strings"
	"sync"
	"constraints"
	"fmt"

	"github.com/jordanorelli/generic/iter"
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
		l.Push(vals[i])
	}
	return l
}

// Empty is true for empty lists
func (l List[T]) Empty() bool {
	return l.head == nil
}

// At treats the list like an array and gets the value at the i'th position in
// the list (zero-indexed)
func (l List[T]) At(i int) T {
	for n, at := l.head, 0; n != nil; n, at = n.next, at+1 {
		if at == i {
			return n.val
		}
	}
	var v T
	return v
}

// Push adds an element to the front of the list
func (l *List[T]) Push(v T) {
	l.head = &node[T]{
		val: v,
		next: l.head,
	}
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

// Head returns the first element of the list. If the list is empty, Head
// returns the zero-value for the type T. This is the same thing as Peek()
func (l List[T]) Head() T {
	if l.head == nil {
		var v T
		return v
	}
	return l.head.val
}

// Tail returns a list which is the original list without its Head element.
// If the original list is an empty list or a list of size 1, Tail is an
// empty list. Note that Tail creates a new list that is backed by the same
// elements as the old list; mutations on the origin list are visible in the
// tail and vice-versa.
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

type _iter[T any] struct {
	n *node[T]
}

func (i _iter[T]) Done() bool { return i.n == nil }

func (i *_iter[T]) Next(dest *T) bool {
	if i.n == nil {
		return false
	}

	*dest = i.n.val
	i.n = i.n.next
	return true
}

func (i *_iter[T]) Iter() iter.Ator[T] { return &_iter[T]{n: i.n} }

func (l List[T]) Iter() iter.Ator[T] { return &_iter[T]{n: l.head} }

func Max[T constraints.Ordered](l List[T]) T {
 	if l.Empty() {
 		var v T
 		return v
 	}
 	
 	v := l.head.val
 	for n := l.head.next; n != nil; n = n.next {
 		if n.val > v {
 			v = n.val
 		}
 	}
 	return v
}

// Map exists as a method to permit chaining in the event that your input
// function maps T -> T. Since methods cannot have type parameters, mapping a
// function that transforms T -> Z is not possible as a method.
func (l List[T]) Map(f func(T) T) List[T] { return Map(l, f) }

// Map applies the input function f to each element of the list l, returning a
// new list containing the values produced by f
func Map[T any, Z any](l List[T], f func(T) Z) List[Z] {
	if l.Empty() {
		var empty List[Z]
		return empty
	}

	mapped := List[Z]{head: &node[Z]{val: f(l.head.val)}}
	last := mapped.head
	for n := l.head.next; n != nil; n = n.next {
		last.next = &node[Z]{val: f(n.val)}
		last = last.next
	}

	return mapped
}

type numbered[T any] struct {
	val T
	i int
}

func waitNClose[T any](wg *sync.WaitGroup, c chan T) {
	wg.Wait()
	close(c)
}

// Run is the same as Map, but is run concurrently. The function f will be run
// for every element of l in its own goroutine. The results of running f on
// each of the inputs will be stored into a new list in an order-preserving
// manner.
func Run[T any, Z any](l List[T], f func(T) Z) List[Z] {
	if l.Empty() {
		var empty List[Z]
		return empty
	}

	// surprise: type declarations are not allowed inside of generic functions
	//
	// type numbered[T any] struct {
	// 	val T
	// 	i int
	// }

	var wg sync.WaitGroup
	c := make(chan numbered[Z])

	i := 0
	for n := l.head; n != nil; n = n.next{
		wg.Add(1)
		go func(v T, i int) {
			defer wg.Done()
			c <- numbered[Z]{val: f(v), i: i}
		}(n.val, i)
		i++
	}

	mem := make([]Z, i)
	go waitNClose(&wg, c)
	for z := range c {
		mem[z.i] = z.val
	}

	var results List[Z]
	for i, _ := range mem {
		results.head = &node[Z]{
			val: mem[i],
			next: results.head,
		}
	}
	return results
}

// Filter applies a predicate function f to each element of the list and
// returns a new list containing the values of the elements that passed the
// predicate
func (l List[T]) Filter(f func(T) bool) List[T] {
	if l.Empty() {
		return List[T]{}
	}

	var passed List[T]
	var last *node[T]
	for n := l.head; n != nil; n = n.next {
		if !f(n.val) {
			continue
		}

		if passed.Empty() {
			passed.head = &node[T]{val: n.val}
			last = passed.head
		} else {
			last.next = &node[T]{val: n.val}
			last = last.next
		}
	}

	return passed
}

type Pair[T any, Z any] struct {
	Left T
	Right Z
}

// Zip takes two lists and joins them to create a list of pairs. It's the same
// as the python zip function, and totally stupid and Pair should not be in
// this package but I'm testing the iterable interfaces and this shows they are
// good, actually
func Zip[T any, Z any](left List[T], right List[Z]) List[Pair[T, Z]] {
	lit, rit := left.Iter(), right.Iter()
	var out List[Pair[T, Z]]

	var next Pair[T, Z]
	for lit.Next(&next.Left) && rit.Next(&next.Right) {
		out.head = &node[Pair[T, Z]]{val: next, next: out.head}
	}
	return out
}
