func twoSum(nums []int, target int) []int {
    // target - num が辞書にあるか
    // 辞書に突っ込む O(n)
    numsMap := make(map[int]int,len(nums)-1)
    
    for i, num := range nums {
        if idx,ok := numsMap[target - num]; ok {
            return []int{i,idx}
        }
        numsMap[num] = i
    }
    return []int{}
}
