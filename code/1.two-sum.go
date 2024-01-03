func twoSum(nums []int, target int) []int {
    m := make(map[int]int,0)
    for i,n := range nums {
        m[n] = i
    }

    for i,v := range nums {
        if k,ok := m[target - v]; ok && i != k {
            return []int{i,k}
        }
    }
    return []int{}
}
