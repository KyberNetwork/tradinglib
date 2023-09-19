package queue

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

type Queue[T any] struct {
	head *Node[T]
	tail *Node[T]
	size uint
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (q *Queue[T]) PushBack(val T) {
	node := &Node[T]{
		value: val,
		next:  nil,
		prev:  nil,
	}

	if q.IsEmpty() {
		q.size++
		q.head = node
		q.tail = node
		return
	}

	q.size++
	node.prev = q.tail
	q.tail.next = node
	q.tail = node
}

func (q *Queue[T]) PushFront(val T) {
	node := &Node[T]{
		value: val,
		next:  nil,
		prev:  nil,
	}

	if q.IsEmpty() {
		q.size++
		q.head = node
		q.tail = node
		return
	}

	q.size++
	node.next = q.head
	q.head.prev = node
	q.head = node
}

func (q *Queue[T]) PopBack() (T, bool) {
	var t T

	if q.IsEmpty() {
		return t, false
	}

	q.size--
	t = q.tail.value
	if q.size == 1 {
		q.tail = nil
		q.head = nil
		return t, true
	}

	if q.tail.prev != nil {
		q.tail.prev.next = nil
	}
	q.tail = q.tail.prev

	return t, true
}

func (q *Queue[T]) PopFront() (T, bool) {
	var t T

	if q.IsEmpty() {
		return t, false
	}

	q.size--
	t = q.head.value
	if q.size == 1 {
		q.tail = nil
		q.head = nil
		return t, true
	}

	if q.head.next != nil {
		q.head.next.prev = nil
	}
	q.head = q.head.next

	return t, true
}

func (q *Queue[T]) PeekBack() (T, bool) {
	var t T

	if q.IsEmpty() {
		return t, false
	}

	t = q.tail.value
	return t, true
}

func (q *Queue[T]) PeekFront() (T, bool) {
	var t T

	if q.IsEmpty() {
		return t, false
	}

	t = q.head.value
	return t, true
}

func (q *Queue[T]) List() []T {
	if q.head == nil {
		return nil
	}

	vals := make([]T, 0, q.size)

	curr := q.head
	for curr != nil {
		vals = append(vals, curr.value)
		curr = curr.next
	}

	return vals
}

func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue[T]) Size() uint {
	return q.size
}
