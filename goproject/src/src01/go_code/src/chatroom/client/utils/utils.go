package utils
import(
	"net"
	"src/chatroom/common/message"
	"fmt"
	"encoding/binary"
	//"errors"
	"encoding/json"
)
type Transfer struct{
	Conn net.Conn
	Buf [8096]byte//缓冲

	
}
func (tran *Transfer)ReadPkg() (mes message.Message,err error){
	//buf:=make([]byte,8096)
	fmt.Println("等待读取数据~~~~")
	//conn.Read 在con没有被关闭情况下，才会阻塞
	//如果有任意一方关闭连接，则不会阻塞，返回
	_,err=tran.Conn.Read(tran.Buf[:4])
	if err!=nil{
		//err = errors.New("red pkg header error")
		return
	}
	//fmt.Println("读到的内容=",buf[:4])
	//根据读到buf[:4] 转成uint32类型
	pkgLen :=binary.BigEndian.Uint32(tran.Buf[:4])
	//根据pkgLen 读取消息内容
	n,err := tran.Conn.Read(tran.Buf[:pkgLen])//从连接里读数据写到buf 切片里
	if n!=int(pkgLen)|| err!=nil{
		//err = errors.New("red pkg body error")
		return 
	}
	//把buf切片反序列化
	//var mes message.Message
	//技术就是一层窗户纸
	err=json.Unmarshal(tran.Buf[:pkgLen],&mes)//注意是指针这里
	if err !=nil{
		fmt.Println("反序列化失败",err)
		return
	}
	return

}
func (tran *Transfer)WritePkg(data []byte) (err error){
	//发送长度给对方
	pkgLen:=uint32(len(data))
	//var buf [4]byte
	binary.BigEndian.PutUint32(tran.Buf[:4],pkgLen)
	n,err:=tran.Conn.Write(tran.Buf[:4])
	if  n!=4 || err!=nil {
		fmt.Println("发生失败")
		return err
	}
	//发送data 本身
	n,err=tran.Conn.Write(data)
	if n!=int(pkgLen) || err != nil{
		fmt.Println("发送失败",err)
		return
	}

	return

}