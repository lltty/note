package stack

import (
	"errors"
	"fmt"
	"strings"
)

type StackItem interface{}

type Stack struct {
	Items  []StackItem
	offset int
	cap    int
}

func NewStack(n int) *Stack {
	return &Stack{
		Items:  make([]StackItem, 0, n),
		offset: 0,
		cap:    n,
	}
}

func (s *Stack) Empty() bool {
	return s.offset == 0
}

func (s *Stack) Fill() bool {
	return s.offset == s.cap
}

func (s *Stack) Push(item StackItem) error {
	if s.Fill() {
		return errors.New("当前栈已经满了")
	}

	s.offset += 1
	s.Items = append(s.Items, item)
	return nil
}

func (s *Stack) ToString() string {
	str := []string{}
	for _, v := range s.Items {
		str = append(str, fmt.Sprintf("%v", v))
	}
	return strings.Join(str, ",")
}

func (s *Stack) Pop() (StackItem, error) {
	if s.Empty() {
		return nil, errors.New("当前栈为空")
	}
	s.offset -= 1
	t := s.Items[s.offset]
	s.Items = s.Items[:s.offset]
	return t, nil
}
