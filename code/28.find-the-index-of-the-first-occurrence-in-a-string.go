func strStr(haystack string, needle string) int {
    end := len(needle)
    for i := 0; i < len(haystack)-end+1; i++ {
        if haystack[i:i+end] == needle {
            return i
        }
    }
    return -1
}
