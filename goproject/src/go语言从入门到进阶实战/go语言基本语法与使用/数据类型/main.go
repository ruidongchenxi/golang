package main

import (
	"fmt"
	"image"
	"image/color"
	"math"
//	"runtime/debug"
)

func main()  {
	const size = 300 //声明一个size常量，值为300
	//根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0,0,size,size))//使用

	for x:=0;x<size;x++{
		for y :=0; y<size;y++ {
			pic.SetGray(x,y,color.Gray{255})

		}
	}
	fmt.Printf("%f\n",math.Pi)//按照默认宽度和精度输出
	fmt.Printf("%.2f\n", math.Pi) //按默认宽度，2位精度输出
}
