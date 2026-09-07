package hanlder
const HelloServiceName= "hanlder/HelloService"
type NewHelloService struct{}

func (s *NewHelloService) Hello(request string, reply *string) error {
	//返回值是通过修改reply 的值
	*reply = "Hello," + request
	return nil
}