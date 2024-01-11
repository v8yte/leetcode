func romanToInt(s string) int {
    romanNums := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    
    ans := romanNums[s[0]]

    if len(s) == 1 {
        return ans
    }

    for i := 1; i < len(s); i++ {
        if romanNums[s[i]] <= romanNums[s[i-1]]{
            ans += romanNums[s[i]]
        }else{
            ans += romanNums[s[i]] - romanNums[s[i-1]]*2
        }
    }
    return ans
}
