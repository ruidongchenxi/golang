package main
import (
	"fmt"
)
func main(){
	//fmt.Println("str := \"c:\\Go\\bin\\go.exe\"")
	coon := `第一行
	第二行
	第三行
	第四行
	\r \n
	      `
    fmt.Println(coon)
	var a byte = 'a'
	fmt.Printf("%d %T \n",a,a)
	var b rune = '你'
	fmt.Printf("%d %T \n",b,b)
}