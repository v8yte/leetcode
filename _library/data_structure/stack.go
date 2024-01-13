package stack

import "container/list"

type Stack struct {
    list *list.List
}

func NewStack() *Stack {
    return &Stack{list: list.New()}
}

func (s *Stack) Push(value any) {
    s.list.PushBack(value)
}

func (s *Stack) Pop() any {
    if element := s.list.Back(); element != nil {
        s.list.Remove(element)
        return element.Value
    }
    return nil
}

func (s *Stack) Top() any {
    if element := s.list.Back(); element != nil {
        return element.Value
    }
    return nil
}

func (s *Stack) Empty() bool {
    return s.list.Len() == 0
}

func (s *Stack) Len() int {
    return s.list.Len()
}
