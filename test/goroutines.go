package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {

	f("direct")

	// 使用 `go f(s)` 在一个 Go 协程中调用这个函数。
	// 这个新的 Go 协程将会并发(concurrently)执行这个函数
	go f("goroutines")

	// 你也可以为匿名函数启动一个 Go 协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 现在这两个 Go 协程在独立的 Go 协程中异步(asynchronously)运行，所以
	// 程序直接运行到这一行。这里的 `Scanln` 代码需要我们
	// 在程序退出前按下任意键结束
	fmt.Scanln()
	fmt.Println("done")
}