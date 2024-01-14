type Solution struct {
    slow int
    fast int
}

func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    s := &Solution{slow: 0, fast: 0}
    for s.fast < len(nums) && s.slow < len(nums) {
        if nums[s.fast] != nums[s.slow] {
            s.slow++
            nums[s.slow] = nums[s.fast]
        }else{
            s.fast++
        }
    }
    return s.slow+1
}
    
