package client_proxy

import (
	"lnhgo/day05/new_hellowoder_test/hanlder"
	"net/rpc"
)

type HelloServiceStub struct {
	*rpc.Client
}
func NewHelloServiveClient(protol,address string) HelloServiceStub{
	conn,err:= rpc.Dial(protol,address)
	if err!=nil{
		panic("链接建立失败")
	}
	return HelloServiceStub{conn}


}

func (h *HelloServiceStub) Hello(request string, reply *string) error {
	err:=h.Call(hanlder.HelloServiceName+".Hello",request,reply)
	if err !=nil{
		return  err
	}
	return nil
}
