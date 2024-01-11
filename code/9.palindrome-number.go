func isPalindrome(x int) bool {
    if x < 0 {
        return false
    }

    X := x
    var y int
    for x > 0 {
        y = 10*y + x%10
        x /= 10
    }
    return y == X
}
