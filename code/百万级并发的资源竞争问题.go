package main

import (
	"fmt"
	"runtime"
	"time"
)

var (testMap = make(map[int]int, 10))


func main() {
	start := time.Now()
	for i := 1; i < 200; i ++ {
		go multiNum(i) // 开启协程, 并发计算阶乘结果
	}

	// 协程需要在 mian 之后完毕
	time.Sleep(time.Second * 5)
	for key, value := range multiNum { // 遍历 map
		fmt.Println("key:", key, "value:", value)
	}
	end := time.Since(start) // Since 返回时间间隔
	fmt.Println("耗时：", end)
}


func multiNum(num int) {
	res :=1
	for i := 1; i <= num; i++ {
		res *= i // 阶乘
	}
	testMap[num] = res // 将结果存入 map
}