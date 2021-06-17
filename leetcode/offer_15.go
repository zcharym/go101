package leetcode

import "math/bits"

func hammingWeight(num uint32) int {
	count := 0
	for i := 0; i < 32; i++ {
		if 1<<i&num > 0 {
			count++
		}
	}
	return count

}

func hammingWeightV2(num uint32) int {
	return bits.OnesCount(uint(num))
}
