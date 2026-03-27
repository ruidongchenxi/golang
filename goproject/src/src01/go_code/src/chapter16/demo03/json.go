package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	//"testing/iotest"

	//"io"
	"os"
)
type Monster struct{
	Nama string
	Age int 
	Skill string
}
func (t *Monster)Store() bool{
	data,err:=json.Marshal(t)
	if err != nil{
		fmt.Printf("错误err=%v\n",err)
		return false
	}
	fild := "d:/Monster.ser"
	err=os.WriteFile(fild,data,0666)
	if err !=nil{
		fmt.Printf("保存出错err=%v",err)
		return false
	}
	return true
}
func (t *Monster)ReStore() bool{
	//读取序列化后字符串
	fild := "d:/Monster.ser"
	data,err:=os.ReadFile(fild)
	if err !=nil{
		fmt.Printf("读取文件错,误err=%v\n",err)
		return false
	}
	err=json.Unmarshal([]byte(data),t)
	if err !=nil{
        fmt.Printf("Unmarshal err=%v\n",err)
    }
	//fmt.Println(t)
	return true

}
func Writer(){
	file,err:= os.OpenFile("d:/a.txt",os.O_WRONLY|os.O_CREATE,0666)
	if err !=nil{
		fmt.Printf("错误err=%v\n",err)
		return
	}
	defer file.Close()
	m := Monster{
		Nama: "孙悟空",
		Age: 1500,
		Skill: "七十二变",
	}
	data,err:=json.Marshal(&m)
	if err !=nil {
		fmt.Printf("序列化失败，err=%V\n",err)
		return
	}
	Writer:=bufio.NewWriter(file)
	Writer.WriteString(string(data))
	Writer.Flush()
}
func Read(){
	//file,err:= os.OpenFile("d:/a.txt",os.O_RDONLY,0666)
	file:="d:/a.txt"

	//defer file.Close()
	// if err !=nil{
	// 	fmt.Printf("文件打开失败err=\n",err)
	// 	return
	// }
	// read:=bufio.NewReader(file)
	t,err :=os.ReadFile(file)
	if err!=nil{
        return
    }
	m := Monster{}
	//str,err:=read.ReadString('}')
	// if err ==io.EOF{
	// 	return
	// }
	
	err=json.Unmarshal([]byte(t),&m)
	if err !=nil{
        fmt.Printf("Unmarshal err=%v\n",err)
    }
	fmt.Println(m)

}
