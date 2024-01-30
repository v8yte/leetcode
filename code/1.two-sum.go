func twoSum(nums []int, target int) []int {
    hashmap := map[int]int{}
    for i,num := range nums {
        if v,ok := hashmap[target-num]; ok{
            return []int{v,i}
        }
        hashmap[num] = i
    }
    return []int{}
}
