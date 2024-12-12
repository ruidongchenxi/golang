package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"math"
	"os"
)

func main()  {
	const size = 300 //声明一个size常量，值为300
	//根据给定大小创建灰度图
	pic := image.NewGray(image.Rect(0,0,size,size))//使用

	for x:=0;x<size;x++{
		s := float64(x) * 2 * math.Pi / size
		y := size/2 - math.Sin(s)*size/2
		pic.SetGray(x,int(y), color.Gray{0})
		/*
		for y :=0; y<size;y++ {
			pic.SetGray(x,y,color.Gray{255})

		}

		 */
		file, err := os.Create("sin.png")
		if err != nil{
			log.Fatal(err)
		}
		png.Encode(file,pic)
		file.Close()
	}

}
