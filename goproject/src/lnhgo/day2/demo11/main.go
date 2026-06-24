package main
import "fmt"
func main(){
	str := []string{"a","b","c","d","e"}
	index:=2
	
	fmt.Println(str[:index],str[index+1:]) 
	str=append(str[:index],str[index+1:]...) //删除指定索引
	fmt.Println(str)
}