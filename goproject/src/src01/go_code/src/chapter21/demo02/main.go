package main
import "fmt"
func SetWay(mymap *[8][7]int,i int,j int) bool{
	//什么情况找到出路
	//myMap[6][5]==2 ；出口
	if mymap[6][5]==2{ 
		return  true
	}else{
		if mymap[i][j]==0{
			//假设这个点是通的；但实际通不通需要探测
			mymap[i][j]=2
			if SetWay(mymap,i+1,j){//下
				return  true
			}else if SetWay(mymap,i,j+1){//右
				return  true
			}else if SetWay(mymap,i-1,j){//上
				return  true
			}else if SetWay(mymap,i,j-1){//左
				return  true
			}else{
				mymap[i][j]=3
				return false
			}


		}else{
			return false
		}
	}


}
func main(){
	//规则
	//如果元素值为1就表示墙
	//2.元素为0 表示路没有走过
	//3. 如果元素2 值，表示通路
	//4.如果元素的值为3，表示走过的路，走不通

	var mymap [8][7]int
	//先把地图最上和最小设置为墙1
	for i:=0;i<7;i++{
		mymap[0][i]=1
		mymap[7][i]=1
		if mymap[i][0]!=1{ 
			mymap[i][0]=1
		}
		if mymap[i][6]!=1{
			mymap[i][6]=1
		}
		
	}
	mymap[3][1]=1
	mymap[3][2]=1
	// mymap[1][2]=1
	// mymap[2][2]=1
	SetWay(&mymap,1,1)
	fmt.Println("探测完毕")
	for i:=0;i<8;i++{
		for j:=0;j<7;j++{
			fmt.Printf("%d ",mymap[i][j])
		}
		fmt.Println()
	}
	//SetWay(&mymap,1,1)
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter21\demo02\main.go
// 探测完毕
// 1 1 1 1 1 1 1 
// 1 2 0 0 0 0 1 
// 1 2 2 2 0 0 1 
// 1 1 1 2 0 0 1 
// 1 0 0 2 0 0 1 
// 1 0 0 2 0 0 1 
// 1 0 0 2 2 2 1 
// 1 1 1 1 1 1 1 