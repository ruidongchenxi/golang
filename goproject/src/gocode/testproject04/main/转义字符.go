package main
import "fmt"
func main(){
	//转义符\n 表示回车换行
	fmt.Println("aaaa\nbbbbb")
	//转义符\b 前移动退格
	fmt.Println("aaassissssss\bbbbbbbb")
	//\r光标回到本行开头后续继续覆盖是输出
	fmt.Println("aaaaaaa\rvvvvv")
	//\t 制表符
	fmt.Println("aaaaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbb")
	fmt.Println("\taaaa")
	//"输出
	fmt.Println("aaaaa\"bbbb\"bbbb")
}