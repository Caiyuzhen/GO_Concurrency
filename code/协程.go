/*
	 Case: 统计 1～2000000 内的素数

		传统的方式: for 循环
		优化的方式: 并发 + 并行 => 在 Go 中就是【协程】 => 【单线程】中的并发, 微线程, 可以开启上万, 其他语言的线程最多开启上千
*/
package main

import (
	"fmt"
	"time"
)

/*
	 Case (要让下面两个程序同时进行):
	 	主线程开启一个 goroutine, 每隔 1s 输出 "执行了 goroutine"
		在主线程中每隔 2s 输出 "执行了主线程" 十次后退出程序
		主线程和 goroutine 同时执行
*/
func main() {
	go runTimes(10) // go 开启协程, 但是主死从随
	for i := 0; i <= 10; i++ {
		fmt.Println("main: ", i, "执行了 goroutine:", 10-i)
		time.Sleep(time.Second * 1) // 👈👈 如果主程序不休眠, runTime 就只会执行一次!
		fmt.Println("            ")
	}
}

func runTimes(times int) int {
	for i := 0; i <= times; i++ {
		fmt.Println("runtimes: ", i, "执行了主线程: ", times-i)
		time.Sleep(time.Second * 1)
	}
	return times
}
