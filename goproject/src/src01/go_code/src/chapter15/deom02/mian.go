package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)
func main(){
	
	//file,err:=os.Open("")
	
	file,err := os.Open("D:/BaiduNetdiskDownload/第5章-基于云原生分布式存储Ceph实现K8S数据持久化/做实验需要的资料/cephfs-pod-1.yaml")
	if err !=nil{
		fmt.Println("open file err =",err)
	}
	defer file.Close()//及时关闭文件
	//创建一个*Reade,默认4096字节
	/*const (
	defaultBufSize = 4096
	)*/
	reade := bufio.NewReader(file)
	for {
		str,err:=reade.ReadString('\n') //读到换行符结束一行
		if err == io.EOF{ //io.EOF,表示文件结尾
			break
		}
		fmt.Print(str) 
	}

}
//执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\deom02\mian.go
// apiVersion: v1
// kind: Pod
// metadata:
//   name: cephfs-pod-1
// spec:
//   containers:
//     - image: nginx
//       name: nginx
//       imagePullPolicy: IfNotPresent
//       volumeMounts:
//       - name: test-v1
//         mountPath: /mnt
//   volumes:
//   - name: test-v1
//     persistentVolumeClaim:
//       claimName: cephfs-pvc