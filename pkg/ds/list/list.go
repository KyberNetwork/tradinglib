package list

type List[T any] []T

func (l List[T]) Filter(pre func(e T) bool) List[T] {
	res := make([]T, 0)
	for _, v := range l {
		if pre(v) {
			res = append(res, v)
		}
	}
	return res
}

func (l List[T]) First() (T, bool) {
	if len(l) > 0 {
		return l[0], true
	}
	var empty T
	return empty, false
}

func (l List[T]) Last() (T, bool) {
	if len(l) > 0 {
		return l[len(l)-1], true
	}
	var empty T
	return empty, false
}

func (l List[T]) Len() int {
	return len(l)
}
