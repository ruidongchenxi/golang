package main
import(
	"fmt"
)
func main(){
	//默认情况下，管道是双向的---->可读可写
	
	//声明为只写
	var intChan1 chan<- int//管道具备<只写性值
    //初始化
	intChan1 = make(chan int,3) 
	//声明只写
	intChan1 <- 20
	//num :=<- intChan1 报错
	fmt.Println("intChan1: ",intChan1)
	//fmt.Println("num",num)
    var intChan2 <-chan int//只读管道
	intChan2 = make(chan int,3)
	//num :=<- intChan2
    if intChan2 != nil {
		num :=<- intChan2
		fmt.Println("num",num)
	}
}