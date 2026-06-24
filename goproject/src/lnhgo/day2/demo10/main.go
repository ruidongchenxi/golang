package main
import "fmt"
func main(){
	const elementCount= 1000
	srcData := make([]int,elementCount)
	for i:=0;i<len(srcData);i++{
		srcData[i]=i
	}
	refData:=srcData
	fmt.Printf("%p\n",&refData)//地址是0xc000008060
	fmt.Printf("%p\n",&srcData)//地址是：0xc000008048

	CopyData:=make([]int,elementCount)
	copy(CopyData,srcData)
	fmt.Printf("%p\n",&CopyData)//0xc000008078	
	srcData[0]=999
	fmt.Println(refData[0]) //999
	fmt.Println(CopyData[0],CopyData[elementCount-1])//
	copy(CopyData,srcData[4:6]) //将srcData局部数据内容4,5复制一份到CopyData[0]CopyData[1]
	for i:=0;i<5;i++{
		fmt.Println(CopyData[i])
	}


}
//执行结果
// PS D:\golang\goproject\src\lnhgo> go run day2\demo10\main.go
// 0xc000008060
// 0xc000008048
// 0xc000008078
// 999
// 0 999
// 4
// 5
// 2
// 3
// 4