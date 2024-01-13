/* スライスでスタック
func isValid(s string) bool {
    brackets := map[rune]rune{
        ']':'[',
        ')':'(',
        '}':'{',
    }

    stack := []rune{}
    for _,c := range s {
        if v,ok := brackets[c]; ok {
            if len(stack) == 0 || v != stack[len(stack)-1] {
                return false
            }
            stack = stack[:len(stack)-1]
        }else{
            stack = append(stack, c)
        }
    }
    return len(stack) == 0
}
*/

func isValid(s string) bool {
    brackets := map[rune]rune{
        ']':'[',
        ')':'(',
        '}':'{',
    }

    stack := NewStack()
    for _,c := range s {
        if v,ok := brackets[c]; ok {
            if stack.Len() == 0 || v != stack.Top() {
                return false
            }
            stack.Pop()
        }else{
            stack.Push(c)
        }
    }
    return stack.Len() == 0
}



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
