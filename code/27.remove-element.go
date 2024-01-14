type Pointer struct {
    slow int
    fast int
}

func removeElement(nums []int, val int) int {
    if len(nums) == 0 {
        return 0
    }
    p := Pointer{
        slow:-1,
        fast:0,
    }

    for p.slow < len(nums) && p.fast < len(nums) {
        if nums[p.fast] != val {
            p.slow++
            nums[p.slow] = nums[p.fast]
        }
        p.fast++
    }
    if p.slow < 0 {
        return 0
    }
    return p.slow+1
}
