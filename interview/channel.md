# channel

- 通道类型值本身是并发安全

```go

// 数据类型为队列

ch1 := make(chan int) // 非缓存通道
ch2 := make(chan int, 3) // 缓冲通道

```