package finderengine

type Comparator[T any] func(a, b T) bool

type MaxHeap[T any] struct {
	data    []T
	compare Comparator[T]
}

func New[T any](cmp Comparator[T]) *MaxHeap[T] {
	return &MaxHeap[T]{
		compare: cmp,
	}
}

func (h *MaxHeap[T]) Push(val T) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap[T]) Pop() (T, bool) {
	var zero T
	if len(h.data) == 0 {
		return zero, false
	}
	top := h.data[0]
	last := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	if len(h.data) > 0 {
		h.data[0] = last
		h.siftDown(0)
	}
	return top, true
}

func (h *MaxHeap[T]) Peek() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}
	return h.data[0], true
}

func (h *MaxHeap[T]) Len() int {
	return len(h.data)
}

func (h *MaxHeap[T]) siftUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if !h.compare(h.data[i], h.data[p]) {
			break
		}
		h.data[i], h.data[p] = h.data[p], h.data[i]
		i = p
	}
}

func (h *MaxHeap[T]) siftDown(i int) {
	n := len(h.data)
	for {
		l, r := 2*i+1, 2*i+2
		largest := i
		if l < n && h.compare(h.data[l], h.data[largest]) {
			largest = l
		}
		if r < n && h.compare(h.data[r], h.data[largest]) {
			largest = r
		}
		if largest == i {
			break
		}
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}
