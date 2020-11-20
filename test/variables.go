package main

import "fmt"

func main() {
	// `var` 声明 1 个或多个变量
	var a = "initial"
	fmt.Println(a)

	// 一次声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// go 将自动推断已经初始化的变量类型
	var d = true
	fmt.Println(d)

	// 声明后没有给对应的初始值，变量会被初始化为 零值。
	var e int
	fmt.Println(e)

	// := 语法是声明并初始化变量的简写
	// 这个例子等于 var f string = "short"
	f := "short"
	fmt.Println(f)
}
