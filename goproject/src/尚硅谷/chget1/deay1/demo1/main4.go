package main

import (
	"fmt"
)

func main() {
	// 有符号整数示例
	var a int = -100      // 默认 int 类型
	var b int8 = 127      // int8 最大值
	var c int16 = -30000  // int16 示例
	var d int32 = 2147483647  // int32 最大值
	var e int64 = -9223372036854775808 // int64 最小值

	// 无符号整数示例
	var f uint = 100      // 默认 uint 类型
	var g uint8 = 255     // uint8 最大值
	var h uint16 = 65000  // uint16 示例
	var i uint32 = 4294967295  // uint32 最大值
	var j uint64 = 18446744073709551615 // uint64 最大值

	fmt.Println("有符号整数:")
	fmt.Println("a (int):", a)
	fmt.Println("b (int8):", b)
	fmt.Println("c (int16):", c)
	fmt.Println("d (int32):", d)
	fmt.Println("e (int64):", e)

	fmt.Println("\n无符号整数:")
	fmt.Println("f (uint):", f)
	fmt.Println("g (uint8):", g)
	fmt.Println("h (uint16):", h)
	fmt.Println("i (uint32):", i)
	fmt.Println("j (uint64):", j)
}