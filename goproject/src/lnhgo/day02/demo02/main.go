package main
import(
	"fmt"
)
func main(){
	var c []string
	c=append(c, "t")
	fmt.Printf("%T\n" ,c)	
	//切片初始化：1从数组直接创建2.使用make，3直接声明{}
	var a1=[4]string{"小狗","小鸡","小鸭","小鹿"}
	//使用make创建
	var b1 = make([]string,5)
	copy(b1,a1[0:2])//左闭右开
	C1:=a1[0:1]
	fmt.Printf("%T\n",C1)
	var b2 []string
	//b2[0]="t"会报错因为没有初始化空间
	b2 = append(b2, "R")//调用函数不会报错
	var b3 []string=make([]string, 6)
	copy(b3,a1[:])
	fmt.Println(b3[:3])//左闭右开
	//合并两个切片
	s1:=[]string{"go","grpc"}
	s2:=[]string{"mysql","es"}
	s1=append(s1, s2...)
	
	s3:=[]string{"kafka","nginx","apache"}
	s1= append(s1, s3[1:]...)
	fmt.Println(s1)

    //删除中间某个元素
	s1=append(s1[:2],s1[3:]...)
	fmt.Println(s1)
	//删除开头某个元素
	s1=s1[1:]
	fmt.Println(s1)
	//删除结尾元素
	s1=s1[:len(s1)-1]
	fmt.Println(s1)
	//复制切片
	s2=s1//拷贝底层数组地址值
	s3=s1[:]//拷贝底层数组地址值
	var s4 []string = make([]string, 3)//拷贝要有空间
	copy(s4,s1)//拷贝底层数组的元素值
	s5:=s1//拷贝底层数组地址值

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println("===============================")
	s1[0]="python"
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println("+++++++++++++++++++++++++++++++++++++")
	fmt.Printf("s1地址值=%p\n",&s1)
	fmt.Printf("s2地址值=%p\n",&s2)
	fmt.Printf("s3地址值=%p\n",&s3)
	fmt.Printf("s4地址值=%p\n",&s4)
	fmt.Printf("s5地址值=%p\n",&s5)
	fmt.Println("+++++++++++++++++++++++++++++++++++++")
	fmt.Printf("s2存储的底层数组地址值=%p\n",s2)
	s2 =append(s2, "rw","w","n","t","c")//只有当append 添加到切片不得不扩张容量的时候，他的底层的数组才会重新拷贝并赋值
	fmt.Printf("s2存储的底层数组地址值=%p\n",s2)
	fmt.Println("+++++++++++++++++++++++++++++++++++++")
	fmt.Printf("s1存储的底层数组地址值=%p\n",s1)
	fmt.Printf("s2存储的底层数组地址值=%p\n",s2)
	fmt.Printf("s3存储的底层数组地址值=%p\n",s3)
	fmt.Printf("s4存储的底层数组地址值=%p\n",s4)
	fmt.Printf("s5存储的底层数组地址值=%p\n",s5)
	//fmt.Printf("s6地址址值=%p\n",&s1)
}