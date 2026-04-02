package main
import(
	"net"
	"src/chatroom/common/message"
	"src/chatroom/service/process2"
	"src/chatroom/service/utils"
	"fmt"
	"io"

)
// 一个结构体
type Processor struct{
	Conn net.Conn
}
func (this *Processor)ServiceProcessMes(mes *message.Message) (err error){
	switch mes.Type {
	case message.LoginMesType:
		//处理登录函数
		up :=&process2.UserProcess{
			Conn: this.Conn,
		}
		err =up.ServiceProcessLogin(mes)
	case message.RegisterMesType:
		//处理注册
	default:
		fmt.Println("消息类型不存在，无法处理")

	}
	return 
}
func (this *Processor)Process2()(err error){
	for {
	//这里将读取数据包，直接封装成一个函数readPkg(),返回Messahe,err
	//等待读取
	//创建transfer 实例完成读包
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
		mes,err:= tf.ReadPkg()
		if err !=nil {
			if err == io.EOF{
				//fmt.Println("客户端关闭连结，服务端也关闭")
				return err
				
			}else{
				fmt.Println("获取连接序列化失败",err)
				return err
				
			}
		}
	//	fmt.Println("mes=",mes)
		err = this.ServiceProcessMes(&mes)
		if err !=nil {
			return err
		
		}
		
	}
	

}