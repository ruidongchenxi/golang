package main
import (
	"context"
	"fmt"
	"lnhgo/protobuf/day3/stream_grpc_test/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)
func main() {
	conn, err := grpc.NewClient(
		"127.0.0.1:50052",
        grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil{
		panic(err)
	}
	defer conn.Close()
	c:= proto.NewGreeterClient(conn)
	res,_:=c.GetStream(context.Background(),&proto.StreamReqData{Data: "慕课网"})
	for {
		a,err:=res.Recv()
		if err != nil{
			fmt.Println(err)
			break
		}
		fmt.Println(a.Data)
	}
	//客户端流模式
	putS,_:=c.PutStream(context.Background())
	i :=0
	for{
		i++
		putS.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("慕课网%d",i),
		})
		time.Sleep(time.Second)
		if i>10{
			break
		}
	}
	wg:=sync.WaitGroup{}
	wg.Add(2)
	allStr,_:=c.AllStream(context.Background())//拿到双向流通道
	go func ()  {
		defer wg.Done()
		for {
			data,_:=allStr.Recv()
			fmt.Println("收到客户端消息:"+data.Data)
		}
	}()
	go func ()  {
		defer wg.Done()
		for{
			_=allStr.Send(&proto.StreamReqData{Data: "我是慕课网"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
}