package main
import(
	"fmt"
	"github.com/redis/go-redis/v9"
)
func main(){
	//通过go 向redis 写数据
	conn,err :=redis.Dial("tcp","127.0.0.1:6379")
	if err !=nil{
		fmt.Printf("链接失败err=",err)
		return
	}
	conn

}