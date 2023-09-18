package queue

// Iterator is not concurrent safe, use with care ^^.
type Iterator[T any] struct {
	queue *Queue[T]
	curr  *Node[T]
	start bool
}

func NewIter[T any](q *Queue[T]) *Iterator[T] {
	return &Iterator[T]{
		queue: q,
		curr:  nil,
	}
}

func (i *Iterator[T]) Val() (T, bool) {
	var t T
	if i.curr == nil {
		return t, false
	}

	return i.curr.value, true
}

func (i *Iterator[T]) Next() bool {
	if !i.start {
		i.start = true
		i.curr = i.queue.head
		return true
	}

	if i.curr == nil || i.curr.next == nil {
		return false
	}

	i.curr = i.curr.next
	return true
}

func (i *Iterator[T]) RemoveCurrent() {
	// 0 node
	if i.curr == nil {
		return
	}

	// 1 node
	if i.queue.head == i.queue.tail {
		i.curr = nil
		i.queue.PopBack()
		return
	}

	tmp := i.curr
	i.curr = i.curr.before
	if tmp == i.queue.head {
		i.queue.PopFront()
		return
	}
	if tmp == i.queue.tail {
		i.queue.PopBack()
		return
	}
	if tmp.before != nil {
		tmp.before.next = tmp.next
	}
	if tmp.next != nil {
		tmp.next.before = tmp.before
	}
}

func (i *Iterator[T]) Reset() {
	i.curr = nil
	i.start = false
}
