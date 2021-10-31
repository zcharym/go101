package leetcode

func moveZeroes(nums []int) {
    var i int
    n := len(nums) - 1
    for ; i < n; i++ {
        if nums[i] != 0 {
            i++
        }
    }
}
