package utils

type Stack[T any] []T

func New[T any]() Stack[T] {
	return make(Stack[T], 0)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) PushN(items []T) {
	for _, item := range items {
		*s = append(*s, item)
	}
}

func (s *Stack[T]) Pop() T {
	if s.IsEmpty() {
		var emptyResult T
		return emptyResult
	}
	index := len(*s) - 1
	pop := (*s)[index]
	*s = (*s)[:index]
	return pop
}

func (s *Stack[T]) PopN(n int) []T {
	if s.IsEmpty() {
		return nil
	}
	index := len(*s) - n
	pop := make([]T, 0)
	for i := index; i < len(*s); i++ {
		pop = append(pop, (*s)[i])
	}
	*s = (*s)[:index]
	return pop
}

func (s *Stack[T]) Peek() T {
	if s.IsEmpty() {
		var emptyResult T
		return emptyResult
	}
	index := len(*s) - 1
	pop := (*s)[index]
	return pop
}
