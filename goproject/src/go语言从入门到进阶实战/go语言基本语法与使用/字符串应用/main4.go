
package main
import (
	"fmt"
	//"strings"
	"bytes"
)
func main(){
	/*
	go 语言和大多数语言一样使用+对字符串拼接操作，非常直观
	但问题来了，好的事物并非完美，简单的东西并非高效。除了加号字符串连接，go 语言中也支持StringBuilder的机制进行高效的字符连接

	*/
	hammer := "吃我一锤"
	sicke := "死吧"
	var stringBuilder bytes.Buffer
	//
	stringBuilder.WriteString(hammer)
	stringBuilder.WriteString(sicke)
	fmt.Println(stringBuilder.String())

}
/*
bytes.Buffer是可以缓冲并可以往里面写入各种字节数组的，字符串也是一种数组，使用WriteString()方法写入
将需要链接的字符串，通过WriteString()方法，写入stringBuilder中，然后通过stringBuilder.String()方法将缓冲转为字符串

*/
