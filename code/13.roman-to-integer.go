func romanToInt(s string) int {
    romanMap := map[byte]int{
        'I':1,
        'V':5,
        'X':10,
        'L':50,
        'C':100,
        'D':500,
        'M':1000,
    }
    var ans int

    for i := 0; i < len(s)-1; i++ {
        if romanMap[s[i]] < romanMap[s[i+1]]{
            ans -= romanMap[s[i]]
        }else{
            ans += romanMap[s[i]]
        }
    }
    
    return ans + romanMap[s[len(s)-1]]
}
