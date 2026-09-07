package server_proxy

import (
	"lnhgo/day05/new_hellowoder_test/hanlder"
	"net/rpc"
)
type HelloService interface{
	Hello(request string,reply *string) error
}
//如果做到解耦-我们关心的是方法，也就是接口

func RegisterHelloService(srv HelloService) error {
	return  rpc.RegisterName(hanlder.HelloServiceName,srv)

}