
package main


import (
	"fmt"
)

func main() {
	// //i 表示层数
	// for i := 0; i<10;i++{
	// 	//j表示打印*
	// 	for j:=0;j<10;j++{
	// 		fmt.Printf("*")
	// 	}
	// 	fmt.Println("")
	// }
	fmt.Println("-----------------------------")
	for i := 1; i<=10;i++{
		for c :=1;c<=10-i;c++{
			fmt.Printf(" ")
		}
	
		//j表示打印*
		for j:=1;j<=i*2-1;j++{
			if j == 1|| j == 2*i-1 || i == 10{// 判断第一个，最后一个；最后一行
			 fmt.Printf("*")
			}else {
				fmt.Printf(" ")
			}
		}
		fmt.Println("")
	}
	fmt.Println("===========================")
	for i :=1;i <=5;i++{
		for k :=1; k<=5-i;k++{
			fmt.Printf(" ")
		}
		for g :=1; g<=2*i-1;g++{
			fmt.Printf("*")
		}
		fmt.Println()
	}
	fmt.Println("------------------------------")
	var h int = 5
	// fmt.Println("输入层数")
	// fmt.Scanln(&h)
	for i:=1;i<=h;i++{ //层数
		//打印左侧空行
		for j:=1; j<=h-i;j++{ 
			fmt.Printf(" ")
		}
		for k:=1; k<=2*i-1;k++{
			if k==1 || k==2*i-1 || i==h{// 第一个打印* 最后一行打印* ;最后一行全打印*
				fmt.Printf("*")
			}else{
				fmt.Printf(" ")
			}
			
		}
		
		fmt.Println()
	}
	
}

// 执行结果
// PS D:\golang\goproject\src\src01\go_code> go run src\chapter04\demo19\main.go
// -----------------------------
//          *
//         * *
//        *   *
//       *     *
//      *       *
//     *         *
//    *           *
//   *             *
//  *               *
// *******************
// ===========================
//     *
//    ***
//   *****
//  *******
// *********
// ------------------------------
//     *
//    * *
//   *   *
//  *     *
// *********