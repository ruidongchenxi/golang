package main
import (
	"encoding/base64"
	"fmt"
)
func main(){
	//需要处理的字符串
	message := "Away from keyboard. https://golang.org/ 尘曦"
	//编码消息
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(message))
	//输出编码完成的信息
	fmt.Println(encodedMessage)
	//解码
	data, err := base64.StdEncoding.DecodeString(encodedMessage)
	if err != nil {
		fmt.Println(err)
	}else{
		//打印解码完成后的信息
		fmt.Println(string(data))
	}
}