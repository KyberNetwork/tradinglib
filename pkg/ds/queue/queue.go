package queue

type Node[T any] struct {
	value  T
	next   *Node[T]
	before *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
	}
}

func (q *Queue[T]) PushBack(val T) {
	node := &Node[T]{
		value:  val,
		next:   nil,
		before: nil,
	}

	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}

	node.before = q.tail
	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) PushFront(val T) {
	node := &Node[T]{
		value:  val,
		next:   nil,
		before: nil,
	}

	if q.head == nil {
		q.head = node
		q.tail = node
		return
	}

	node.next = q.head
	q.head.before = node
	q.head = node
}

func (q *Queue[T]) PopBack() (T, bool) {
	var t T

	// 0 node
	if q.tail == nil {
		return t, false
	}

	t = q.tail.value
	// only 1 node
	if q.tail == q.head {
		q.tail = nil
		q.head = nil
		return t, true
	}

	if q.tail.before != nil {
		q.tail.before.next = nil
	}
	q.tail = q.tail.before

	return t, true
}

func (q *Queue[T]) PopFront() (T, bool) {
	var t T

	// 0 node
	if q.head == nil {
		return t, false
	}

	t = q.head.value
	// only 1 node
	if q.tail == q.head {
		q.tail = nil
		q.head = nil
		return t, true
	}

	if q.head.next != nil {
		q.head.next.before = nil
	}
	q.head = q.head.next

	return t, true
}

func (q *Queue[T]) List() []T {
	if q.head == nil {
		return nil
	}

	var vals []T

	curr := q.head
	for curr != nil {
		vals = append(vals, curr.value)
		curr = curr.next
	}

	return vals
}
