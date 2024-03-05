package main

import (
	"fmt"
	"runtime"
	"time"
)

var num int = 1


// 主线程
func main() {
	for i := 1; i <= 1000000; i++ { // 百万并发
		go Run(i) // 开启协程
	}

	// 设置 CPU 核心数
	runtime.GOMAXPROCS(8) // 设置 CPU 的最大运行数

	fmt.Println("当前电脑的 CPU 核心有:", runtime.NumCPU()) // 获取 CPU 核心数

	for i := 1; i <= 10; i ++ {
		fmt.Println("main: ", i, "执行了 goroutine:", 10-i)
		time.Sleep(time.Second * 1)
	}
}


func Run(times int) int {
	for i := 1; i <= times; i ++ {
		fmt.Println("runtimes: ", i, "执行了 time 线程: ", times - i)
		fmt.Println("Num: ", num)
		time.Sleep(time.Second * 1)
	}
	num ++ // 计算并发数
	return times
}