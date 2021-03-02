package queue

import (
	"errors"
	"fmt"
	"strings"
)

type QueueItem interface{}

type Queue struct {
	Items  []QueueItem
	offset int
	cap    int
}

func NewQueue(n int) *Queue {
	return &Queue{
		Items:  make([]QueueItem, 0, n),
		offset: 0,
		cap:    n,
	}
}

func (s *Queue) Empty() bool {
	return s.offset == 0
}

func (s *Queue) Fill() bool {
	return s.offset == s.cap
}

func (s *Queue) Push(item QueueItem) error {
	if s.Fill() {
		return errors.New("当前队列已经满了")
	}

	s.offset += 1
	s.Items = append(s.Items, item)
	return nil
}

func (s *Queue) ToString() string {
	str := []string{}
	for _, v := range s.Items {
		str = append(str, fmt.Sprintf("%v", v))
	}
	return strings.Join(str, ",")
}

func (s *Queue) Pop() (QueueItem, error) {
	if s.Empty() {
		return nil, errors.New("当前队列为空")
	}
	t := s.Items[0]
	if s.offset == 1 {
		s.Items = s.Items[0:0]
	} else {
		s.Items = s.Items[1:s.offset]
	}
	s.offset -= 1
	return t, nil
}
