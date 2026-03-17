package main
import (
	"fmt"
	//"slices"
)
type Person struct{
	x int
	y int 
}
type Rect struct{
	leftUP,rightDown Person
}
type Rect1 struct{
	leftUP,rightDown *Person
}
func main(){
	r1 :=Rect{Person{1,2},Person{3,5}}
	fmt.Printf("r1.leftUP.x 地址=%p r1.leftUP.y 地址=%p r1.rightDown.x 地址=%p r1.rightDown.y 地址=%p\n",
	&r1.leftUP.x,&r1.leftUP.y,&r1.rightDown.x,&r1.rightDown.y)
	//r2 有两个*Point类型，这两个*Point类型的本身地址是连续的，但是他们指向的地址不一定是连续的,也有可能是连续的
	r2 :=Rect1{&Person{10,20},&Person{20,30}}
	fmt.Printf("r2.leftUP.地址=%p r2.rightDown 地址=%p\n",&r2.leftUP,&r2.rightDown)
	fmt.Printf("r2.leftUP.x 地址=%p r2.leftUP.y 地址=%p r2.rightDown.x 地址=%p r2.rightDown.y 地址=%p\n",
	&r2.leftUP.x,&r2.leftUP.y,&r2.rightDown.x,&r2.rightDown.y)

}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter09\deom03\main.go
// r1.leftUP.x 地址=0xc000016220 r1.leftUP.y 地址=0xc000016228 r1.rightDown.x 地址=0xc000016230 r1.rightDown.y 地址=0xc000016238
// r2.leftUP.地址=0xc000026070 r2.rightDown 地址=0xc000026078
// r2.leftUP.x 地址=0xc00000a0d0 r2.leftUP.y 地址=0xc00000a0d8 r2.rightDown.x 地址=0xc00000a0e0 r2.rightDown.y 地址=0xc00000a0e8