package main
import (
	"github.com/redis/go-redis/v9"
	//"fmt"
)
func main(){
	//通过go 向redis 写数据 
	conn:=redis.NewClient(&redis.Options{ Addr: "127.0.0.1:6379", })
	conn
 }
