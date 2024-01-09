package main

type Stack[T any] []T

func (s *Stack[T]) Push(x T) {
	*s = append(*s, x)
}

func (s *Stack[T]) Pop() T {
	pop := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return pop
}

func (s *Stack[T]) Peek() T {
	return (*s)[len(*s)-1]
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

// 写个快排熟悉一下代码吧

func main() {

}
