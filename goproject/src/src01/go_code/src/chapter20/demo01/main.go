package main
import (
	"fmt"
)
type  ValNode struct {
	row int
	col int
	val int
}

func main(){
	var chessMap [11][11]int
	chessMap [1][2] = 1
	chessMap [2][3] = 2
	for _,v:=range chessMap{
		for _,o:= range v{
			fmt.Printf("%v\t",o)
		}
		fmt.Println()
	}
	//转成稀疏数组
	var sparseArr []ValNode
	t := ValNode{
					row: 11,
					col: 11,
					val: 2,
				}
	sparseArr=append(sparseArr,t)
	//标准稀疏数组，应该还有一个表示记录原始的二维数据的规模（行和列，以及默认值）
	for i,v:=range chessMap{
		for r,o:= range v{
			if o !=0{
				t = ValNode{
					row: i,
					col: r,
					val: o,
				}
				sparseArr=append(sparseArr,t)
			}
		}
		//fmt.Println()
	}
	//fmt.Println()
	for i,v:= range sparseArr{
		fmt.Printf("%d:%d %d %d\n",i,v.row,v.col,v.val)
	}
	var chessMap2 [11][11]int
	for i,valNade := range sparseArr{
		if i!=0{
			chessMap2[valNade.row][valNade.col]= valNade.val
		}
	}
	for _,v:=range chessMap2{
		for _,o:= range v{
			fmt.Printf("%v\t",o)
		}
		fmt.Println()
	}
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter20\demo01\main.go
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       1       0       0       0       0       0       0       0       0
// 0       0       0       2       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0:11 11 2
// 1:1 2 1
// 2:2 3 2
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       1       0       0       0       0       0       0       0       0
// 0       0       0       2       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0
// 0       0       0       0       0       0       0       0       0       0       0