package main
import (
	"time"
)
//定义time.Duration的别名为MyDuration
type MyDuration = time.Duration
func (m MyDuration) EasySet(a string) {

}
func main(){
	 
}
/*
编译报错
# command-line-arguments                                                                                                
D:\golang\goproject\src\go语言从入门到进阶实战\go语言基本语法与使用\别名类型\main0.go:7:7: cannot define new methods on non-local type time.Duration    
编译器: 不能一个非本地方的类型time.Duration上定义新方法，非本地方法指的就是使用time.Duration的代码所在的包，也就是main包。因为time.Duration是在time包中定义的，在main包中使用。time.Duration包与main包不在同一个包中，因为不能为不在同一个包中的类型定义方法
解决此问题
修改为 type MyDuration time.Duration,也就是将MyDuration 从别名改为类型
将MyDuration 的别名定义放在time包里
*/