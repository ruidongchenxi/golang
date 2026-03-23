package main
import(
	"os"
	"fmt"
)
func main(){
	//大开文件
	//file 对象
	//
	file,err := os.Open("D:/BaiduNetdiskDownload/第5章-基于云原生分布式存储Ceph实现K8S数据持久化/做实验需要的资料/cephfs-pod-1.yaml")
	if err !=nil{
		fmt.Println("open file err =",err)
	}else{
		fmt.Printf("文件内容=%v",file)
	}
	defer file.Close()
	if err != nil {
		fmt.Println("close file err=",err)
	}
	
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\deom01\main.go
// 文件内容=&{0xc000070a08}
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\deom01\main.go
// 文件内容=&{0xc000070a08}
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\deom01\main.go
// open file err = open D:/课程资料/crictl0.yaml: The system cannot find the file specified.
// close file err= invalid argument