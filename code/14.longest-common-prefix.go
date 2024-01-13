// Horizontal scanning
/*
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    cp := strs[0]
    for i := 1; i < len(strs); i++ {
        idx := min(len(cp),len(strs[i]))
        cp = cp[:idx]
        for j := 0; j < idx; j++ {
            if cp[j] != strs[i][j] {
                cp = cp[:j]
                break
            }
        }
    }
    return cp
}
*/

// Vertical scanning
func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    var builder strings.Builder
    var current string
    for i := 0; i < len(strs[0]); i++ {
        current = string(strs[0][i])
        for j := 1; j < len(strs); j++ {
            if i > len(strs[j])-1 || strs[0][i] != strs[j][i] {
                return builder.String()
            }
        }
        builder.WriteString(current)
    }
    return builder.String()
}
