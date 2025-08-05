package main

import (
	"fmt"
	"math/big"
)

func main() {
	// 棋盘格子总数
	const totalSquares = 

	// 创建一个切片用于存储每个格子的值
	board := make([]int, totalSquares)

	// 填充棋盘值
	for i := 0; i < totalSquares; i++ {
		board[i] = 1 << i // 2 的 i 次方
	}

	// 打印棋盘值
	for i, value := range board {
		fmt.Printf("格子 %d 的值是: %d\n", i+1, value)
	}
}
