package stack

// Stack is a generic stack.
type Stack[T any] struct {
	elems []T
}

// New returns a new instance of Stack.
// hint: hint for stack size.
func New[T any](hint int) *Stack[T] {
	if hint < 0 {
		hint = 0
	}

	return &Stack[T]{elems: make([]T, 0, hint)}
}

// Empty returns true if the stack is empty.
func (s *Stack[T]) Empty() bool {
	return len(s.elems) == 0
}

// Len returns number of elements in the stack.
func (s *Stack[T]) Len() int {
	return len(s.elems)
}

// Peek views the top element of the stack.
func (s *Stack[T]) Peek() (elem T, ok bool) {
	if s.Empty() {
		return
	}

	return s.elems[len(s.elems)-1], true
}

// Push pushes a new element into the stack.
func (s *Stack[T]) Push(elem T) {
	s.elems = append(s.elems, elem)
}

// Pop pops the top element from the stack.
func (s *Stack[T]) Pop() (elem T, ok bool) {
	if s.Empty() {
		return
	}

	n := s.Len()
	elem = s.elems[n-1]
	s.elems = s.elems[:n-1]

	return elem, true
}
