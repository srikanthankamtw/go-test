package generics

// Stack that only supports integer values

type StackOfInts struct {
	values []int
}

func (s *StackOfInts) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfInts) Push(value int) {
	s.values = append(s.values, value)
}

func (s *StackOfInts) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	}
	lastIndex := len(s.values) - 1
	lastElement := s.values[lastIndex]
	s.values = s.values[:lastIndex]
	return lastElement, true
}

// Stack that only supports string values

type StackOfStrings struct {
	values []string
}

func (s *StackOfStrings) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *StackOfStrings) Push(value string) {
	s.values = append(s.values, value)
}

func (s *StackOfStrings) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}
	lastIndex := len(s.values) - 1
	lastElement := s.values[lastIndex]
	s.values = s.values[:lastIndex]
	return lastElement, true
}

// Generic way of implementing Stack

type Stack[T any] struct {
	values []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var defaultValue T
		return defaultValue, false
	}
	lastIndex := len(s.values) - 1
	lastElement := s.values[lastIndex]
	s.values = s.values[:lastIndex]
	return lastElement, true
}
