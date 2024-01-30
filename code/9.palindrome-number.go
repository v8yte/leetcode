func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }
    
    X := x
    var y int
    for X > 0 {
        y = y*10 + X%10
        X /= 10
    }
    
    return x == y
}
