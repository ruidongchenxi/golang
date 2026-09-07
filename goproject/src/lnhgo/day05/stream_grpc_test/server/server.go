package main

import (
	"fmt"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"

	"lnhgo/day05/stream_grpc_test/proto"
)

const PORT = ":50052"

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server)GetStream(req *proto.StreamReqDate, res grpc.ServerStreamingServer[proto.StreamResData]) error{
	i :=0
	for{
		i++
		_=res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v",time.Now().Unix()),
		})
		time.Sleep(time.Second)
		if i>10{
			break
		}

	}
	return nil
}
func (s *server)PostStream(cliStr grpc.ClientStreamingServer[proto.StreamReqDate, proto.StreamResData]) error{
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
func (s *server)AllStream(allStr grpc.BidiStreamingServer[proto.StreamReqDate, proto.StreamResData]) error{
	// go allStr.Recv()
	// allStr.Send()\
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			data,_:=allStr.Recv()
			fmt.Println("收到客户端消息:"+data.Data)
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
//func (s *server)mustEmbedUnimplementedGreeterServer(){}
func main(){
	list,err:=net.Listen("tcp",PORT)
	if err !=nil{
		panic("端口启动失败")
	}
	s:=grpc.NewServer()
	proto.RegisterGreeterServer(s,&server{})
	err=s.Serve(list)
	if err !=nil{
		panic("启动失败")
	}

}