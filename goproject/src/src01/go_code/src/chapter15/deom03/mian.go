package main

import (
	"bufio"
	"fmt"
	"os"
	//"strings"
)
func main(){
	//file,err:=os.Open("")
	
	// file :="D:/BaiduNetdiskDownload/第5章-基于云原生分布式存储Ceph实现K8S数据持久化/做实验需要的资料/cephfs-pvc.yaml"//只适合小文件
	// t,err :=os.ReadFile(file)
	// if err!=nil{
	// 	return
	// }
	// fmt.Printf("%v",string(t))//t 默认是byte 切片所有这里需要使用强转转为string类型
	s,err:= os.Open("D:/BaiduNetdiskDownload/第5章-基于云原生分布式存储Ceph实现K8S数据持久化/做实验需要的资料/cephfs-pvc.yaml")
	if err !=nil{
		return
	}
	defer s.Close()
	t :=bufio.NewScanner(s)
	for t.Scan() { //此方法返回布尔值
		line := t.Text() // 当前行
		fmt.Println(line)
	}
	if err := t.Err(); err != nil {
		fmt.Println("读取错误:", err)
	} 
}
// 执行结果
// PS D:\golang\goproject\src\src01\go_code\src> go run chapter15\deom03\mian.go
// kind: PersistentVolumeClaim
// apiVersion: v1
// metadata:
//   name: cephfs-pvc
// spec:
//   accessModes:
//     - ReadWriteMany
//   volumeName: cephfs-pv
//   resources:
//     requests:
//       storage: 1Gi