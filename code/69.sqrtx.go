func mySqrt(x int) int {
    var ans int
    for ans * ans <= x {
        ans++
    }
    return ans - 1
}
