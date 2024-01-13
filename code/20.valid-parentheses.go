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
