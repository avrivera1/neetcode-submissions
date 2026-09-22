
func hasDuplicate(nums []int) bool {
    //return true if any value in an array appears more than once
    m := make(map[int]bool)

    for i := 0; i < len(nums); i++ {
        if m[nums[i]] {
            return true
        }
        m[nums[i]] = true
    }
    return false
}
