package main
import(
	"fmt"
)
func main(){
	var c byte//只要存英文字符
	c = 'a'+1 //也是字符可以汉字
	fmt.Printf("c=%c\n",c)
	var c2 rune
	c2 = '熙'
	fmt.Printf("c2=%c\n",c2)
	var name string 
	name = "亚马逊"
	fmt.Println(name)
}