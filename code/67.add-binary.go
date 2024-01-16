func addBinary(a string, b string) string {
    return decimalToBinary(binaryToDecimal(a) + binaryToDecimal(b))
}

func binaryToDecimal(x string) int{
    var decimal int
    for i,v := range x { 
        value,_ := strconv.Atoi(string(v))
        decimal += int(math.Pow(2,float64(len(x)-i-1))) * value
    }
    return decimal
}

func decimalToBinary(x int) string{
    if x == 0 {
        return "0"
    }

    var builder strings.Builder
    for x > 0 {
        builder.WriteString(strconv.Itoa(x%2))
        x /= 2
    }
    return reverse(builder.String())
}

func reverse(s string) string{
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
