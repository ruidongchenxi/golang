package main
import (
	"fmt"
)
func main(){
	//将int 定义为Weapon类型。就像枚举类型其实本质就是int 一样，当然某些情况下，如果需要int32 和int64 的枚举，也可以
	type Weapon int
	/*
	一个const 声明内的每一行常量声明，将会自动套用前面的iota格式，并自动增加。这种模式有点类似于电子表格中的单元格自动填写。只需建立好单元格之间的变化关系
	
	*/
	const (
		//将arrow常量的类型标识为Weapon,这样标识后，const 下方的常量可以是默认类型的，默认时，默认使用前面指定的类型作为常量类型，该行使用iota进行常量自动生成
		//iota起始值为0，一般情况下也是建议枚举从0开始，让每个枚举类型都有一个空值，方便业务和逻辑的使用
		Arrow Weapon = iota //开始生成枚举值，默认0
		Shuriken
		SniperRifle
		Rifle
		Blower
	)

	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)
	var weapon Weapon = Blower
	fmt.Println(weapon)
	/*
	当然。iota不仅只生成每次增加1的枚举值，我们还可以利用iota来做一些强大的枚举常量值生成器，下面的代码可以方便生成标志位常量

	*/
	const (
		//iota使用了一个移位操作，每次将上一次的值像左移一位，以做出每一位的常量值
		FlagNone = 1 << iota
		//将3个枚举按照常量输出，分别是2、4、8, 都是将1每次左移一位的结果
		FlagRed
		FlagGreen
		FlagBlue
	) 
	fmt.Printf("%d %d %d\n", FlagRed,FlagGreen,FlagBlue)
	fmt.Printf("%b %b %b\n", FlagRed,FlagGreen,FlagBlue)//将枚举值按二进制格式输出，可以清晰看到每一位的变化
}