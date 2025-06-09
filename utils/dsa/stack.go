package dsa

type Stack[T any] struct {
	keys []T
}

func (stack *Stack[T]) Push(key T) {
	stack.keys = append(stack.keys, key)
}

func (stack *Stack[T]) Top() T {
	x := stack.keys[len(stack.keys)-1]
	return x
}

func (stack *Stack[T]) Pop() T {
	var x T
	x, stack.keys = stack.keys[len(stack.keys)-1], stack.keys[:len(stack.keys)-1]
	return x
}

func (stack *Stack[T]) IsEmpty() bool {
	return len(stack.keys) == 0
}
