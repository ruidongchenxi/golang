package main

import (
	"fmt"

	//"golang.org/x/tools/go/analysis/passes/hostport"
)
type MyWriter interface{
	Write(string) error

}
type MyCloser interface{
	Close() error
}
type writerCloser struct{
	 MyWriter//一个接口实现
}
type fileWriter struct{}
func (f *fileWriter)Write(a string) error{
	fmt.Println("打开文件")
	return nil
}
type database struct{
	host string
}
func (f *database)Write(a string) error{
	fmt.Println("链接数据库")
	return nil
}
// func (wc *writerCloser)Write(a string) error{
// 	fmt.Println("write  string")
// 	return nil
// }
func (wc *writerCloser)Close() error{
	fmt.Println("Close")
	return  nil
}
func main(){
	var md MyWriter = &writerCloser{
	//	&fileWriter{},//注入接口实现
		 &database{},
	}
	md.Write("a")

}