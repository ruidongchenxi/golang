package main
import (
	"fmt"
	//"os"
	"flag"
)
func  main(){
//定义几个变量，用于接收
	var  user string
	var  pwd string
	var  host string
	var  prot int
	//&user用于接收用户命令行输入的-u,后面参数值
	//u 就是-u 这个参数
	//"" 表示默认值
	//"用户名默认为空" 对参数说明
	flag.StringVar(&user,"u","","用户名默认为空")
	flag.StringVar(&pwd,"pwd","","密码默认为空")
	flag.StringVar(&host,"h","localhost","主机默认为localhost")
	flag.IntVar(&prot,"port",3306,"默认3306")
	//这里有一个非常重要操作，转换，必须调用方法
	flag.Parse()
	fmt.Printf("user=%v;pwd=%v;host=%v;prot=%v\n",user,pwd,host,prot)
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\deom13\mian.go -u chenxi -pwd 123.com -h 127.0.0.1 -port 3306         
// user=chenxi;pwd=123.com;host=127.0.0.1;prot=3306