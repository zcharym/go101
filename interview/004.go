package interview

import (
	"fmt"
)

/**
分析关于该函数可能的执行流程
- 不打印任何东西
- 打印"The data value is: 0"
- 打印"The data value is: 1"
*/
func solution004() {
	var data int
	go func() {
		data++
	}()
	if data == 1 {
		fmt.Printf("The data value is: %d", data)
	}
}
