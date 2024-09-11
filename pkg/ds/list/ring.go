package list

const (
	defaultInitSize = 4
	resizeFactor    = 1.5
)

// Ring is a circular, resizable ring style buffer.
type Ring[T any] struct {
	items     []T
	writeIdx  int
	tailIdx   int
	itemCount int
}

// NewRing create a new ring with init size.
func NewRing[T any](initSize int) *Ring[T] {
	if initSize < 0 {
		initSize = 0
	}
	return &Ring[T]{
		items: make([]T, initSize),
	}
}

func (r *Ring[T]) Len() int {
	return r.itemCount
}

// Append add item at right position of the buffer, expand it if needed.
func (r *Ring[T]) Append(item T) {
	currentSize := len(r.items)
	if (r.itemCount) == currentSize {
		if currentSize == 0 {
			r.expand(defaultInitSize)
		} else {
			r.expand(int(float64(currentSize) * resizeFactor))
		}
	}
	r.items[r.writeIdx] = item
	r.writeIdx = (r.writeIdx + 1) % len(r.items)
	r.itemCount++
}

// Expire remove n item at the left position of the ring.
func (r *Ring[T]) Expire(n int) {
	if n <= 0 {
		return
	}
	if n > r.itemCount {
		n = r.itemCount
	}
	r.tailIdx = (r.tailIdx + n) % len(r.items)
	r.itemCount -= n
}

// ExpireCond is keep removing item at the left position if predicate return true,
// it stops at the first false value.
func (r *Ring[T]) ExpireCond(pre func(e T) bool) {
	removeCount := 0
	r.ScanAsc(func(item T) bool {
		if pre(item) {
			removeCount++
			return true
		}
		return false
	})
	r.Expire(removeCount)
}

// ScanAsc iterator items from left to right direction.
func (r *Ring[T]) ScanAsc(fn func(e T) bool) {
	for i := 0; i < r.itemCount; i++ {
		if !fn(r.items[(i+r.tailIdx)%len(r.items)]) {
			break
		}
	}
}

// ScanReverse iterator items from right to left direction.
func (r *Ring[T]) ScanReverse(fn func(e T) bool) {
	for i := 0; i < r.itemCount; i++ {
		if !fn(r.items[(r.writeIdx-i-1+len(r.items))%len(r.items)]) {
			break
		}
	}
}

// Filter collect matched item into a slice.
func (r *Ring[T]) Filter(fn func(e T) bool) List[T] {
	res := make(List[T], 0, r.itemCount)
	r.ScanAsc(func(e T) bool {
		if fn(e) {
			res = append(res, e)
		}
		return true
	})
	return res
}

func (r *Ring[T]) expand(newSize int) {
	if newSize <= len(r.items) {
		return
	}
	newSlice := make([]T, newSize)
	if r.tailIdx < r.writeIdx {
		copy(newSlice, r.items[r.tailIdx:r.writeIdx])
	} else {
		p1 := len(r.items) - r.tailIdx
		copy(newSlice, r.items[r.tailIdx:])
		copy(newSlice[p1:], r.items[0:r.writeIdx])
	}
	r.items = newSlice
	r.writeIdx = r.itemCount
	r.tailIdx = 0
}
