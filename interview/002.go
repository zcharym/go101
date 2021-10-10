package interview

import (
	"fmt"
	"time"
)

/**
Q：求打印结果
A：没有任何打印结果:for语句执行完毕时，其中的十个go函数还没有获得运行的机会
*/
func solution002() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 500)
		go func() {
			fmt.Println(i)
		}()
	}
}

/**
Q：求打印结果
A：可能为打印十个10
*/
func solution002v2() {
	num := 10
	sign := make(chan struct{}, num)
	for i := 0; i < num; i++ {
		go func() {
			fmt.Println(i)
			sign <- struct{}{}
		}()
	}

	for j := 0; j < num; j++ {
		<-sign
	}
}

/**
Q：求打印结果
A：乱序打印0-9
*/
func solution002v3() {
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Millisecond * 5000)
}
