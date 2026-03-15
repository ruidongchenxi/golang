package main
import (
	"fmt"
)
func main(){
	var name string
	names := [4]string{"白眉鹰王","金毛狮王","青翼蝠王","紫衫龙王"}
	fmt.Println("输入名字")
	fmt.Scanln(&name)
	//顺序查找：第一种方式
	// for i:=0; i<len(names);i++{
	// 	if name == names[i]{
	// 		fmt.Printf("找到%v,下标%v\n",name,i)
	// 		break
	// 	}else if i == len(names)-1{
	// 		fmt.Println("没有找到",name)
	// 	}

	// }
	// for i,v:=range names{
	// 	if name ==v {
	// 		fmt.Printf("找到%v,下标%v\n",name,i)
	// 		break
	// 	} else if i == len(names)-1{
	// 		fmt.Println("没有找到",name)
	// 	}
	// }
	index := -1
	for i,v:=range names{
		if name== v{
			index= i
			break
		}
	}
	if index>=0{
		fmt.Printf("找到了%v，下标是%v",names[index],index)
	}
}
// 执行方式
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 杨逍
// 没有找到 杨逍
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 紫衫龙王
// 找到紫衫龙王,下标3
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 张无忌
// 没有找到 张无忌
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 青翼蝠王
// 找到青翼蝠王,下标2
// 没有找到 青翼蝠王
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 青翼蝠王
// 找到青翼蝠王,下标2
// 没有找到 青翼蝠王
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 2
// 没有找到 2
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 没有找到 
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 青翼蝠王
// 找到青翼蝠王,下标2
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 紫衫龙王
// 找到紫衫龙王,下标3
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 紫衫龙王
// 找到了紫衫龙王，下标是3
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter07\demo02\main.go
// 输入名字
// 白眉鹰王
// 找到了白眉鹰王，下标是0