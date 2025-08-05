package main
import "fmt"
func main(){
	//连续转义字符
	fmt.Println("aaaaaaa\nbbbbbbbbb")
	//\b 退格
	fmt.Println("aaaaaaaaaaa\bbbbbbbbbb")
//\r 回车   bbbbbbbbbaa
	fmt.Println("aaaaaaaaaaa\rbbbbbbbbb")
	//\t
	fmt.Println("aaaaaaaaaaa\tbbbbbbbbb")
	fmt.Println("aaaaaaaaa\tbbbbbbbbb")
}