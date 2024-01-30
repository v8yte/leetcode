func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a
    }

    indexB := len(b)-1
    ans := make([]byte,len(a))
    var shifter,sum byte
    for i := len(a)-1; i >= 0; i-- {
        sum = shifter + a[i] - '0'
        if indexB >= 0 {
            sum += b[indexB] - '0'
            indexB--
        }
        ans[i] = sum%2 + '0'
        shifter = sum/2
    }
    if shifter == 0 {
        return string(ans)
    }
    return "1" + string(ans)
}
