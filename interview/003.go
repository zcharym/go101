package interview

import (
	"fmt"
)

func solution003() {
	nums := []int{1, 2, 3, 4, 5, 6}
	for i := range nums {
		// 当只有一个迭代变量时，我们只能拿到迭代对应元素值的索引值
		fmt.Println(i)
		if i == 3 {
			nums[i] |= i // 4 ｜ 3 结果为7
		}
	}
	fmt.Println(nums) // []int{1,2,3,7,5,6}
}
