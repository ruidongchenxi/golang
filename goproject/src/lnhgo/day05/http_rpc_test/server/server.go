package main

import (
	//"net"
	//"encoding/json"
	//"net"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	//"structs"
)
type HelloService struct{

}
func (s *HelloService)Hello(requst string,reply *string) error{
	*reply = "hello,"+requst
	return nil
}

func main() {
	_=rpc.RegisterName("HelloService",&HelloService{})  //注册rpc
	//1.实例化一个server
	http.HandleFunc("/jsonrpc",func(w http.ResponseWriter, r *http.Request){
		var conn io.ReadWriteCloser=struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser:  r.Body,
			Writer: w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})
	http.ListenAndServe(":1234",nil)

	
}