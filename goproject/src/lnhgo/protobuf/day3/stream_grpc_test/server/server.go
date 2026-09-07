package main
import (
	"fmt"
	"lnhgo/protobuf/day3/stream_grpc_test/proto"
	"net"
	"sync"
	"time"
	"google.golang.org/grpc"
)
const POST = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}
func (s *server)GetStream(req *proto.StreamReqData, res grpc.ServerStreamingServer[proto.StreamResData]) error{
	i:=0
	for{
		i++
		_ =res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v",time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i>10{
			break
		}
	}
	return nil
}
func (s *server)PutStream(cliStr proto.Greeter_PutStreamServer) error{
    for {
        if a,err:=cliStr.Recv();err!=nil{
            fmt.Println(err)
            break
        }else{
            fmt.Println(a.Data)
        }
    }
    return nil
}
func (s *server)AllStream(allStr proto.Greeter_AllStreamServer) error{
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func ()  {
		defer wg.Done()
		for {
			data,_:=allStr.Recv()
			fmt.Println("收到客户端消息"+data.Data)
		}
	}()
	go func ()  {
		defer wg.Done()
		for {
			_=allStr.Send(&proto.StreamResData{Data: "你好我是服务器"})
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	return  nil
}
func main(){
	lis,err:=net.Listen("tcp",POST)
	if err!=nil{
		panic(err)
	}
	s:=grpc.NewServer()
	proto.RegisterGreeterServer(s,&server{})
	err= s.Serve(lis)
	if err!=nil{
		panic(err)
	}
}